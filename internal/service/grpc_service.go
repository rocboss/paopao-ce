// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"google.golang.org/grpc"
)

type baseGRPCService struct {
	baseService

	server *grpcServer
}

func (s *baseGRPCService) registerServer(srv Service, h func(s *grpc.Server)) {
	h(s.server.server)
	s.server.addService(srv)
}

func (s *baseGRPCService) OnStart() error {
	// do nothing default
	return nil
}

func (s *baseGRPCService) OnStop() error {
	// do nothing default
	return nil
}
