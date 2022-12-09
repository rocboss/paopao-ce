// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/gin-gonic/gin"
)

type httpService interface {
	Service

	registerRoute(e *gin.Engine)
}

type baseHttpService struct {
	baseService

	server *httpServer
}

func (s *baseHttpService) Init() error {
	if s.server.status() != httpServerStarted {
		s.registerRoute(s.server.e)
	}
	return nil
}

func (s *baseHttpService) Start() error {
	if err := s.server.start(); err != nil {
		return err
	}
	return nil
}

func (s *baseHttpService) Stop() error {
	return s.server.stop()
}

func (s *baseHttpService) registerRoute(e *gin.Engine) {
	// default empty
}

func newBaseHttpService(s *httpServer) httpService {
	return &baseHttpService{
		server: s,
	}
}
