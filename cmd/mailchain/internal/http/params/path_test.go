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

package params

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mailchain/mailchain/internal/mail"
	"github.com/stretchr/testify/assert"
)

func TestPathMessageID(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    mail.ID
		wantErr bool
	}{
		{
			"success",
			args{
				func() *http.Request {
					req := httptest.NewRequest("GET", "/message_id", nil)
					req = mux.SetURLVars(req, map[string]string{
						"message_id": "002c47eca011e32b52c71005ad8a8f75e1b44c92c99fd12e43bccfe571e3c2d13d2e9a826a550f5ff63b247af471",
					})
					return req
				}(),
			},
			[]byte{0x0, 0x2c, 0x47, 0xec, 0xa0, 0x11, 0xe3, 0x2b, 0x52, 0xc7, 0x10, 0x5, 0xad, 0x8a, 0x8f, 0x75, 0xe1, 0xb4, 0x4c, 0x92, 0xc9, 0x9f, 0xd1, 0x2e, 0x43, 0xbc, 0xcf, 0xe5, 0x71, 0xe3, 0xc2, 0xd1, 0x3d, 0x2e, 0x9a, 0x82, 0x6a, 0x55, 0xf, 0x5f, 0xf6, 0x3b, 0x24, 0x7a, 0xf4, 0x71},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathMessageID(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathMessageID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(tt.want, got) {
				t.Errorf("PathMessageID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathNetwork(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{
				func() *http.Request {
					req := httptest.NewRequest("GET", "/", nil)
					req = mux.SetURLVars(req, map[string]string{
						"network": "mainnet",
					})
					return req
				}(),
			},
			"mainnet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PathNetwork(tt.args.r)
			if !assert.Equal(tt.want, got) {
				t.Errorf("PathNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathAddress(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		r *http.Request
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
				func() *http.Request {
					req := httptest.NewRequest("GET", "/", nil)
					req = mux.SetURLVars(req, map[string]string{
						"address": "5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761",
					})
					return req
				}(),
			},
			[]byte{0xee, 0xd6, 0xf4, 0x9d, 0x11, 0x7e, 0xab, 0x95, 0xc8, 0xab, 0x8b, 0x71, 0xba, 0xe2, 0xcd, 0xd1, 0xe5, 0x64, 0xa7, 0x61},
			false,
		},
		{
			"success",
			args{
				func() *http.Request {
					req := httptest.NewRequest("GET", "/", nil)
					req = mux.SetURLVars(req, map[string]string{
						"address": "",
					})
					return req
				}(),
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathAddress(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(tt.want, got) {
				t.Errorf("PathAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
