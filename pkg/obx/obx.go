// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// package obx contain some help function for OpenObserve.
package obx

import (
	"bytes"
	"net/http"

	hx "github.com/rocboss/paopao-ce/pkg/http"
)

var (
	_ OpenObserveClient = (*obxClient)(nil)
)

const (
	_userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"
)

// OpenObserveClient OpenObserve client interface
type OpenObserveClient interface {
	LogJson(data []byte)
}

// Config confiugre used for create a OpenObserveClient instance
type Config struct {
	Host         string
	User         string
	Password     string
	Organization string
	Stream       string
	UserAgent    string
	Secure       bool
}

type obxClient struct {
	endpoint  string
	user      string
	password  string
	userAgent string
	respFn    hx.ResponseFn
	client    hx.AsyncClient
}

func (c *Config) Endpoint() string {
	schema := "http"
	if c.Secure {
		schema = "https"
	}
	return schema + "://" + c.Host + "/api/" + c.Organization + "/" + c.Stream + "/_json"
}

func (s *obxClient) LogJson(data []byte) {
	req, err := http.NewRequest("POST", s.endpoint, bytes.NewReader(data))
	if err != nil {
		s.respFn(nil, nil, err)
	}
	req.SetBasicAuth(s.user, s.password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", s.userAgent)
	s.client.Do(req, s.respFn)
}

// NewClient create OpenObserve client instance
func NewClient(conf *Config, acc *hx.AsyncClientConf, fn hx.ResponseFn) OpenObserveClient {
	userAgent := _userAgent
	if conf.UserAgent != "" {
		userAgent = conf.UserAgent
	}
	return &obxClient{
		endpoint:  conf.Endpoint(),
		user:      conf.User,
		password:  conf.Password,
		userAgent: userAgent,
		respFn:    fn,
		client:    hx.NewAsyncClient(http.DefaultClient, acc),
	}
}
