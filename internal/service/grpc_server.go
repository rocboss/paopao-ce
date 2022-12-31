// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"net"

	"google.golang.org/grpc"
)

var (
	_ server = (*grpcServer)(nil)
)

// grpcServer wraper for grpc.Server
type grpcServer struct {
	*baseServer

	listener net.Listener
	server   *grpc.Server
}

func (s *grpcServer) start() error {
	s.Lock()
	if s.serverStatus == _statusServerStarted || s.serverStatus == _statusServerStoped {
		return nil
	}
	oldStatus := s.serverStatus
	s.serverStatus = _statusServerStarted
	s.Unlock()

	if err := s.server.Serve(s.listener); err != nil {
		s.Lock()
		s.serverStatus = oldStatus
		s.Unlock()

		return err
	}
	return nil
}

func (s *grpcServer) stop() error {
	s.Lock()
	defer s.Unlock()

	if s.serverStatus == _statusServerStoped || s.serverStatus == _statusServerInitilized {
		return nil
	}
	s.server.Stop()
	s.serverStatus = _statusServerStoped
	return nil
}
