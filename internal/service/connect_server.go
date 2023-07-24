// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
)

var (
	_ server = (*connectServer)(nil)
)

// connectServer wraper for connect
type connectServer struct {
	*baseServer

	isTLS       bool
	certFile    string
	keyFile     string
	handlerOpts []connect.HandlerOption
	server      *http.Server
	mux         *http.ServeMux
}

func (s *connectServer) start() error {
	s.server.Handler = s.mux
	if s.isTLS {
		return s.server.ListenAndServeTLS(s.certFile, s.keyFile)
	} else {
		return s.server.ListenAndServe()
	}
}

func (s *connectServer) stop() error {
	return s.server.Shutdown(context.Background())
}

func (s *connectServer) register(path string, handler http.Handler) {
	s.mux.Handle(path, handler)
}

func defaultConnectServer(addr string) *connectServer {
	return &connectServer{
		baseServer: newBaseServe(),
		server: &http.Server{
			Addr: addr,
		},
		mux: &http.ServeMux{},
	}
}
