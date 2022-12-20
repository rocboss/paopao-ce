// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

var (
	httpServers = make(map[string]*httpServer)
	grpcServers = make(map[string]*grpcServer)
)

const (
	_statusServerInitilized uint8 = iota + 1
	_statusServerStarted
	_statusServerStoped
)

func httpServerFrom(addr string, newServer func() *httpServer) *httpServer {
	s, exist := httpServers[addr]
	if exist {
		return s
	}
	s = newServer()
	s.serverStatus = _statusServerInitilized
	httpServers[addr] = s
	return s
}

func grpcServerFrom(addr string, newServer func() *grpcServer) *grpcServer {
	s, exist := grpcServers[addr]
	if exist {
		return s
	}
	s = newServer()
	s.serverStatus = _statusServerInitilized
	grpcServers[addr] = s
	return s
}
