// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import "sync"

var (
	httpServers = newServerPool[*httpServer]()
	grpcServers = newServerPool[*grpcServer]()
)

const (
	_statusServerUnknow uint8 = iota
	_statusServerInitilized
	_statusServerStarted
	_statusServerStoped
)

type server interface {
	status() uint8
	setStatus(uint8)
	start() error
	stop() error
}

type serverPool[T server] struct {
	servers map[string]T
}

type baseServer struct {
	sync.RWMutex
	serverStatus uint8
}

func (p *serverPool[T]) from(addr string, newServer func() T) T {
	s, exist := p.servers[addr]
	if exist {
		return s
	}
	s = newServer()
	s.setStatus(_statusServerInitilized)
	p.servers[addr] = s
	return s
}

func (s *baseServer) setStatus(status uint8) {
	s.RLock()
	defer s.RUnlock()

	s.serverStatus = status
}

func (s *baseServer) status() uint8 {
	s.RLock()
	defer s.RUnlock()

	return s.serverStatus
}

func newServerPool[T server]() *serverPool[T] {
	return &serverPool[T]{
		servers: make(map[string]T),
	}
}

func newBaseServe() *baseServer {
	return &baseServer{
		RWMutex:      sync.RWMutex{},
		serverStatus: _statusServerUnknow,
	}
}
