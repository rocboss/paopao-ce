// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package http

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	_ AsyncClient = (*wormClient)(nil)
)

const (
	_minRequestInCh    = 10
	_minRequestInTmpCh = 10
	_minWorker         = 5
)

// ResponseFn a function used handle the response of http.Client.Do
type ResponseFn = func(req *http.Request, resp *http.Response, err error)

// AsyncClient asynchronous client interface
type AsyncClient interface {
	Do(req *http.Request, fn ResponseFn)
}

// AsyncClientConf client configure used to create an AsynClient instance
type AsyncClientConf struct {
	MinWorker          int
	MaxRequestInCh     int
	MaxRequestInTempCh int
	MaxTickCount       int
	TickWaitTime       time.Duration
}

type requestItem struct {
	request *http.Request
	fn      ResponseFn
}

type wormClient struct {
	client        *http.Client
	requestCh     chan *requestItem // 正式工 缓存通道
	requestTempCh chan *requestItem // 临时工 缓存通道
	maxTickCount  int
	tickWaitTime  time.Duration
}

func (s *wormClient) Do(req *http.Request, fn ResponseFn) {
	item := &requestItem{req, fn}
	select {
	case s.requestCh <- item:
		logrus.Debugln("simepleClient.Do send request item by requestCh chan")
	default:
		select {
		case s.requestTempCh <- item:
			logrus.Debugln("simepleClient.Do send request item by requestTempCh chan")
		default:
			go func() {
				s.do(item)
				// watch requestTempCh to continue do work if needed.
				// cancel loop if no item had watched in s.maxCyle * s.maxWaitTime.
				for count := 0; count < s.maxTickCount; count++ {
					select {
					case item := <-s.requestTempCh:
						// reset count to continue do work
						count = 0
						s.do(item)
					default:
						// sleeping to wait request item pass over to do work
						time.Sleep(s.tickWaitTime)
					}
				}
			}()
		}
	}
}

func (s *wormClient) starDotWork() {
	for item := range s.requestCh {
		s.do(item)
	}
}

func (s *wormClient) do(req *requestItem) {
	resp, err := s.client.Do(req.request)
	req.fn(req.request, resp, err)
}

// NewAsyncClient create an AsyncClient instance
func NewAsyncClient(client *http.Client, conf *AsyncClientConf) AsyncClient {
	maxRequestInCh := _minRequestInCh
	maxRequestInTempCh := _minRequestInTmpCh
	if conf.MaxRequestInCh > _minRequestInCh {
		maxRequestInCh = conf.MaxRequestInCh
	}
	if conf.MaxRequestInTempCh > _minRequestInTmpCh {
		maxRequestInTempCh = conf.MaxRequestInTempCh
	}
	wc := &wormClient{
		client:        client,
		requestCh:     make(chan *requestItem, maxRequestInCh),
		requestTempCh: make(chan *requestItem, maxRequestInTempCh),
	}
	numWorker := conf.MinWorker
	if numWorker < _minWorker {
		numWorker = _minWorker
	}
	logrus.Debugf("use %d backend worker to do the http request", numWorker)
	// 启动 do work 正式工
	for ; numWorker > 0; numWorker-- {
		go wc.starDotWork()
	}
	return wc
}
