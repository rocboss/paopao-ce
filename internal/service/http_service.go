// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/gin-gonic/gin"
)

type baseHttpService struct {
	baseService

	server *httpServer
}

func (s *baseHttpService) registerRoute(h func(e *gin.Engine)) {
	if s.server.status() != httpServerStarted {
		h(s.server.e)
	}
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
