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
	return s.server.Serve(s.listener)
}

func (s *grpcServer) stop() error {
	s.server.Stop()
	return nil
}
