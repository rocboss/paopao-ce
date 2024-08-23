// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package http

import (
	"net/http"
	"time"

	gp "github.com/alimy/tryst/pool"
)

var (
	_ AsyncClient = (*wormClient)(nil)
)

const (
	_minRequestBuf     = 10
	_minRequestTempBuf = 10
	_minWorker         = 5
)

// ResponseFn a function used handle the response of http.Client.Do
type ResponseFn = gp.ResponseFn[*http.Request, *http.Response]

// AsyncClient asynchronous client interface
type AsyncClient interface {
	Do(req *http.Request, fn ResponseFn)
}

// AsyncClientConf client configure used to create an AsynClient instance
type AsyncClientConf struct {
	MinWorker         int
	MaxRequestBuf     int
	MaxRequestTempBuf int
	MaxIdleTime       time.Duration
}

type wormClient struct {
	pool gp.GoroutinePool[*http.Request, *http.Response]
}

func (s *wormClient) Do(req *http.Request, fn ResponseFn) {
	s.pool.Do(req, fn)
}

// NewAsyncClient create an AsyncClient instance
func NewAsyncClient(client *http.Client, conf *AsyncClientConf) AsyncClient {
	minWorker := _minWorker
	maxRequestBuf := _minRequestBuf
	maxRequestTempBuf := _minRequestTempBuf
	if conf.MaxRequestBuf > _minRequestBuf {
		maxRequestBuf = conf.MaxRequestBuf
	}
	if conf.MaxRequestTempBuf > _minRequestTempBuf {
		maxRequestTempBuf = conf.MaxRequestTempBuf
	}
	if conf.MinWorker > _minWorker {
		minWorker = conf.MinWorker
	}
	return &wormClient{
		pool: gp.NewGoroutinePool(func(req *http.Request) (*http.Response, error) {
			return client.Do(req)
		},
			gp.WithMinWorker(minWorker),
			gp.WithMaxRequestBuf(maxRequestBuf),
			gp.WithMaxRequestTempBuf(maxRequestTempBuf),
			gp.WithMaxIdelTime(conf.MaxIdleTime),
		),
	}
}
