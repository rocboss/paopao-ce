// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"net"
	"sync"

	"google.golang.org/grpc"
)

// grpcServer wraper for grpc.Server
type grpcServer struct {
	sync.RWMutex

	listener     net.Listener
	server       *grpc.Server
	serverStatus uint8
}

func (s *grpcServer) status() uint8 {
	s.RLock()
	defer s.RUnlock()

	return s.serverStatus
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
