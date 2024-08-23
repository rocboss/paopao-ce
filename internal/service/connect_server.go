// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	hx "github.com/rocboss/paopao-ce/pkg/http"
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
	mux         connectMux
}

type connectMux interface {
	http.Handler
	Handle(string, http.Handler)
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

func defaultConnectServer(addr string) (s *connectServer) {
	s = &connectServer{
		baseServer: newBaseServe(),
		server: &http.Server{
			Addr: addr,
		},
		mux: &http.ServeMux{},
	}
	// TODO: custom value from config
	var useConnectMux bool
	if useConnectMux {
		s.mux = hx.NewConnectMux()
	}
	return
}
