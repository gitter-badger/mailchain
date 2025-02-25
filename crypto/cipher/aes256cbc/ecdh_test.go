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

package aes256cbc

import (
	"encoding/hex"
	"log"
	"math/big"
	"testing"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/mailchain/mailchain/crypto"
	"github.com/mailchain/mailchain/crypto/secp256k1"
	"github.com/mailchain/mailchain/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestGenerateMacKeyAndEncryptionKey(t *testing.T) {
	assert := assert.New(t)
	secret, err := hex.DecodeString("04695325aac70f9f9ebe676248ebbfefa87b3eff16117559d2a0953d0e695be6")
	if err != nil {
		log.Fatal(err)
	}
	macKey, encryptionKey := generateMacKeyAndEncryptionKey(secret)

	assert.Equal("2cea25760305bdb3194057646bc46dc2eeee4890b711741c0b525454ac7c5ea8", hex.EncodeToString(macKey))
	assert.Equal("af0ad81e7d9194721d6c26f6c1f2a2b7fd06e2c99c4f5deefe59fb93936c981e", hex.EncodeToString(encryptionKey))
}

func TestGenerateIV(t *testing.T) {
	assert := assert.New(t)
	iv, err := generateIV()
	if err != nil {
		log.Fatal(err)
	}
	assert.Len(iv, 16)
}

func TestGenerateMac(t *testing.T) {
	assert := assert.New(t)
	macKey := testutil.MustHexDecodeString("2cea25760305bdb3194057646bc46dc2eeee4890b711741c0b525454ac7c5ea8")
	iv := testutil.MustHexDecodeString("05050505050505050505050505050505")
	cipherText := testutil.MustHexDecodeString("2ec66aac453ff543f47830d4b8cbc68d9965bf7c6bb69724fd4de26d41001256dfa6f7f0b3956ce21d4717caf75b0c2ad753852f216df6cfbcda4911619c5fc34798a19f81adff902c1ad906ab0edaec")
	tmpEphemeralPrivateKey, err := ethcrypto.HexToECDSA("0404040404040404040404040404040404040404040404040404040404040404")
	if err != nil {
		log.Fatal(err)
	}
	ephemeralPrivateKey := ecies.ImportECDSA(tmpEphemeralPrivateKey)
	actual, err := generateMac(macKey, iv, ephemeralPrivateKey.PublicKey, cipherText)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal("4367ae8a54b65f99e4f2fd315ba65bf85e1138967a7bea451faf80f75cdf3404", hex.EncodeToString(actual))
}

func TestEncryptDecrypt(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name                string
		recipientPublicKey  crypto.PublicKey
		recipientPrivateKey crypto.PrivateKey
		data                []byte
		err                 error
	}{
		{"to-sofia-short-text",
			testutil.SofiaPublicKey,
			testutil.SofiaPrivateKey,
			[]byte("Hi Sofia"),
			nil,
		}, {"to-sofia-medium-text",
			testutil.SofiaPublicKey,
			testutil.SofiaPrivateKey,
			[]byte("Hi Sofia, this is a little bit of a longer message to make sure there are no problems"),
			nil,
		},
		{"to-charlotte-short-text",
			testutil.CharlottePublicKey,
			testutil.CharlottePrivateKey,
			[]byte("Hi Charlotte"),
			nil,
		}, {"to-charlotte-medium-text",
			testutil.CharlottePublicKey,
			testutil.CharlottePrivateKey,
			[]byte("Hi Charlotte, this is a little bit of a longer message to make sure there are no problems"),
			nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := Encrypt(tc.recipientPublicKey, tc.data)
			assert.Equal(tc.err, err)
			assert.NotNil(encrypted)
			decrypter := Decrypter{&tc.recipientPrivateKey}

			decrypted, err := decrypter.Decrypt(encrypted)
			assert.Equal(tc.err, err)
			assert.Equal(tc.data, []byte(decrypted))
		})
	}
}

func Test_deriveSharedSecret(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		pub     *ecies.PublicKey
		private *ecies.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"success",
			args{
				func() *ecies.PublicKey {
					tp, ok := testutil.CharlottePublicKey.(secp256k1.PublicKey)
					if !ok {
						t.Error("failed to cast")
					}
					pub, err := tp.ECIES()
					if err != nil {
						t.Error(err)
					}
					return pub
				}(),
				func() *ecies.PrivateKey {
					pk, ok := testutil.SofiaPrivateKey.(*secp256k1.PrivateKey)
					if !ok {
						t.Error("Can not cast")
					}
					return pk.ECIES()
				}(),
			},
			testutil.MustHexDecodeString("b6bdfade23178272425d25774a7d0d388fbef9480893fcc3646accc123eacc47"),
			false,
		},
		{
			"err-scalar-mult",
			args{
				func() *ecies.PublicKey {
					tp, ok := testutil.CharlottePublicKey.(secp256k1.PublicKey)
					if !ok {
						t.Error("failed to cast")
					}
					pub, err := tp.ECIES()
					if err != nil {
						t.Error(err)
					}
					pub.X = big.NewInt(0)
					pub.Y = big.NewInt(0)
					return pub
				}(),
				func() *ecies.PrivateKey {
					pk, ok := testutil.SofiaPrivateKey.(*secp256k1.PrivateKey)
					if !ok {
						t.Error("Can not cast")
					}
					p := pk.ECIES()
					p.D = big.NewInt(0)
					return p
				}(),
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deriveSharedSecret(tt.args.pub, tt.args.private)
			if (err != nil) != tt.wantErr {
				t.Errorf("deriveSharedSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(tt.want, got) {
				t.Errorf("deriveSharedSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
