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

func (s *baseGRPCService) registerServer(h func(s *grpc.Server)) {
	if s.server.status() != _statusServerStarted {
		h(s.server.server)
	}
}

func (s *baseGRPCService) OnStart() error {
	if err := s.server.start(); err != nil {
		return err
	}
	return nil
}

func (s *baseGRPCService) OnStop() error {
	return s.server.stop()
}
