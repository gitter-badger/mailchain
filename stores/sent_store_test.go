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

package stores

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/mailchain/mailchain/internal/mail"
	"github.com/mailchain/mailchain/internal/testutil"
	"github.com/pkg/errors"
)

func Test_responseAsError(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"no-err",
			args{
				&http.Response{
					StatusCode: http.StatusCreated,
				},
			},
			false,
		},
		{
			"err-valid-response",
			args{
				&http.Response{
					StatusCode: http.StatusConflict,
					Body:       ioutil.NopCloser(strings.NewReader("{\"code\": 409, \"message\": \"conflict\"}")),
				},
			},
			true,
		},
		{
			"err-invalid-response",
			args{
				&http.Response{
					StatusCode: http.StatusConflict,
					Body:       ioutil.NopCloser(strings.NewReader("\"code\": 409")),
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := responseAsError(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("responseAsError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewSentStore(t *testing.T) {
	tests := []struct {
		name       string
		wantDomain string
	}{
		{
			"success",
			"https://mcx.mx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSentStore()
			if !reflect.DeepEqual(got.domain, tt.wantDomain) {
				t.Errorf("NewSentStore() = %v, want %v", got, tt.wantDomain)
			}
		})
	}
}

func TestSentStore_PutMessage(t *testing.T) {
	type fields struct {
		newRequest func(method string, url string, body io.Reader) (*http.Request, error)
		doRequest  func(req *http.Request) (*http.Response, error)
	}
	type args struct {
		messageID mail.ID
		msg       []byte
		headers   map[string]string
	}
	tests := []struct {
		name    string
		server  *httptest.Server
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"success",
			httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Location", "https://mcx.mx/mesasgeID-hash")
					w.WriteHeader(http.StatusCreated)
				}),
			),
			fields{
				newRequest: http.NewRequest,
				doRequest: func() func(req *http.Request) (*http.Response, error) {
					c := http.Client{
						Timeout: 1 * time.Second,
					}
					return c.Do
				}(),
			},
			args{
				testutil.MustHexDecodeString("002c47eca011e32b52c71005ad8a8f75e1b44c92c99fd12e43bccfe571e3c2d13d2e9a826a550f5ff63b247af471"),
				[]byte("body"),
				nil,
			},
			"https://mcx.mx/mesasgeID-hash",
			false,
		},
		{
			"err-new-request",
			httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Location", "https://mcx.mx/mesasgeID-hash")
					w.WriteHeader(http.StatusCreated)
				}),
			),
			fields{
				newRequest: func() func(method string, url string, body io.Reader) (*http.Request, error) {
					return func(method string, url string, body io.Reader) (*http.Request, error) {
						return nil, errors.Errorf("failed to create request")
					}
				}(),
				doRequest: func() func(req *http.Request) (*http.Response, error) {
					c := http.Client{
						Timeout: 1 * time.Second,
					}
					return c.Do
				}(),
			},
			args{
				testutil.MustHexDecodeString("002c47eca011e32b52c71005ad8a8f75e1b44c92c99fd12e43bccfe571e3c2d13d2e9a826a550f5ff63b247af471"),
				[]byte("body"),
				nil,
			},
			"",
			true,
		},
		{
			"err-do-request",
			httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Location", "https://mcx.mx/mesasgeID-hash")
					w.WriteHeader(http.StatusCreated)
				}),
			),
			fields{
				newRequest: http.NewRequest,
				doRequest: func() func(req *http.Request) (*http.Response, error) {
					return func(req *http.Request) (*http.Response, error) {
						return nil, errors.Errorf("do request failed")
					}
				}(),
			},
			args{
				testutil.MustHexDecodeString("002c47eca011e32b52c71005ad8a8f75e1b44c92c99fd12e43bccfe571e3c2d13d2e9a826a550f5ff63b247af471"),
				[]byte("body"),
				nil,
			},
			"",
			true,
		},
		{
			"err-in-request",
			httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Location", "https://mcx.mx/mesasgeID-hash")
					w.WriteHeader(http.StatusConflict)
				}),
			),
			fields{
				newRequest: http.NewRequest,
				doRequest: func() func(req *http.Request) (*http.Response, error) {
					c := http.Client{
						Timeout: 1 * time.Second,
					}
					return c.Do
				}(),
			},
			args{
				testutil.MustHexDecodeString("002c47eca011e32b52c71005ad8a8f75e1b44c92c99fd12e43bccfe571e3c2d13d2e9a826a550f5ff63b247af471"),
				[]byte("body"),
				nil,
			},
			"",
			true,
		},
		{
			"err-missing-location",
			httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusCreated)
				}),
			),
			fields{
				newRequest: http.NewRequest,
				doRequest: func() func(req *http.Request) (*http.Response, error) {
					c := http.Client{
						Timeout: 1 * time.Second,
					}
					return c.Do
				}(),
			},
			args{
				testutil.MustHexDecodeString("002c47eca011e32b52c71005ad8a8f75e1b44c92c99fd12e43bccfe571e3c2d13d2e9a826a550f5ff63b247af471"),
				[]byte("body"),
				nil,
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SentStore{
				domain:     tt.server.URL,
				newRequest: tt.fields.newRequest,
				doRequest:  tt.fields.doRequest,
			}
			got, err := s.PutMessage(tt.args.messageID, tt.args.msg, tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("SentStore.PutMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SentStore.PutMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSentStore_Key(t *testing.T) {
	type fields struct {
		domain     string
		newRequest func(method string, url string, body io.Reader) (*http.Request, error)
		doRequest  func(req *http.Request) (*http.Response, error)
	}
	type args struct {
		messageID mail.ID
		msg       []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"success",
			fields{
				"",
				nil,
				nil,
			},
			args{
				testutil.MustHexDecodeString("002c47eca011e32b52c71005ad8a8f75e1b44c92c99fd12e43bccfe571e3c2d13d2e9a826a550f5ff63b247af471"),
				[]byte("message"),
			},
			"002c47eca011e32b52c71005ad8a8f75e1b44c92c99fd12e43bccfe571e3c2d13d2e9a826a550f5ff63b247af471-220455078214",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SentStore{
				domain:     tt.fields.domain,
				newRequest: tt.fields.newRequest,
				doRequest:  tt.fields.doRequest,
			}
			if got := s.Key(tt.args.messageID, tt.args.msg); got != tt.want {
				t.Errorf("SentStore.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
