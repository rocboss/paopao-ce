// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"net/http"

	"github.com/bufbuild/connect-go"
)

type baseConnectService struct {
	baseService

	server *connectServer
}

func (s *baseConnectService) registerServer(srv Service, h func(func(string, http.Handler), ...connect.HandlerOption), opts ...connect.HandlerOption) {
	h(s.server.register, append(opts, s.server.handlerOpts...)...)
	s.server.addService(srv)
}

func (s *baseConnectService) OnStart() error {
	// do nothing default
	return nil
}

func (s *baseConnectService) OnStop() error {
	// do nothing default
	return nil
}
