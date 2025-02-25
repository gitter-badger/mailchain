// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mailbox

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/proto"
	"github.com/mailchain/mailchain/crypto"
	"github.com/mailchain/mailchain/crypto/cipher/aes256cbc"
	"github.com/mailchain/mailchain/internal/encoding"
	"github.com/mailchain/mailchain/internal/mail"
	"github.com/mailchain/mailchain/internal/mail/rfc2822"
	"github.com/mailchain/mailchain/internal/mailbox/signer"
	"github.com/mailchain/mailchain/sender"
	"github.com/mailchain/mailchain/stores"
	"github.com/pkg/errors"
)

type SendMessageFunc func(ctx context.Context, network string, msg *mail.Message, pubkey crypto.PublicKey,
	sender sender.Message, sent stores.Sent, signer signer.Signer) error

// SendMessageFunc performs all the actions required to send a message.
// - Create a hash of encoded message
// - Encrypt message
// - Store sent message
// - Encrypt message location
// - Create transaction data with encrypted location and message hash
// - Send transaction
func SendMessage() SendMessageFunc {
	return sendMessage(defaultEncryptLocation, defaultEncryptMailMessage, defaultPrefixedBytes)
}

func sendMessage(encryptLocation func(pk crypto.PublicKey, location string) ([]byte, error),
	encryptMailMessage func(pk crypto.PublicKey, encodedMsg []byte) ([]byte, error),
	prefixedBytes func(data proto.Message) ([]byte, error)) SendMessageFunc {
	return func(ctx context.Context,
		network string,
		msg *mail.Message,
		pubkey crypto.PublicKey,
		sender sender.Message,
		sent stores.Sent,
		signer signer.Signer) error {
		encodedMsg, err := rfc2822.EncodeNewMessage(msg)
		if err != nil {
			return errors.WithMessage(err, "could not encode message")
		}
		msgHash := crypto.CreateMessageHash(encodedMsg)

		encrypted, err := encryptMailMessage(pubkey, encodedMsg)
		if err != nil {
			return errors.WithMessage(err, "could not encrypt mail message")
		}

		location, err := sent.PutMessage(msg.ID, encrypted, nil)
		if err != nil {
			return errors.WithStack(err)
		}
		encryptedLocation, err := encryptLocation(pubkey, location)
		if err != nil {
			return errors.WithMessage(err, "could not encrypt transaction data")
		}

		data := &mail.Data{
			EncryptedLocation: encryptedLocation,
			Hash:              msgHash,
		}
		encodedData, err := prefixedBytes(data)
		if err != nil {
			return errors.WithMessage(err, "could not encode transaction data")
		}
		transactonData := append(encoding.DataPrefix(), encodedData...)
		//TODO: should not use common to parse address
		to := common.FromHex(msg.Headers.To.ChainAddress)
		from := common.FromHex(msg.Headers.From.ChainAddress)
		if err := sender.Send(ctx, network, to, from, transactonData, signer, nil); err != nil {
			return errors.WithMessage(err, "could not send transaction")
		}

		return nil
	}
}

// encryptLocation is encrypted with supplied public key and location string
func defaultEncryptLocation(pk crypto.PublicKey, location string) ([]byte, error) {
	// TODO: encryptLocation hard coded to aes256cbc
	return aes256cbc.Encrypt(pk, []byte(location))
}

// encryptMailMessage is encrypted with supplied public key and location string
func defaultEncryptMailMessage(pk crypto.PublicKey, encodedMsg []byte) ([]byte, error) {
	// TODO: encryptMailMessage hard coded to aes256cbc
	return aes256cbc.Encrypt(pk, encodedMsg)
}

func defaultPrefixedBytes(data proto.Message) ([]byte, error) {
	protoData, err := proto.Marshal(data)
	if err != nil {
		return nil, errors.WithMessage(err, "could not marshal data")
	}

	prefixedProto := make([]byte, len(protoData)+1)
	prefixedProto[0] = encoding.Protobuf
	copy(prefixedProto[1:], protoData)

	return prefixedProto, nil
}
