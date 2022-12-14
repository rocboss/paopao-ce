// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/servants/web/routers"
)

type oldWebService struct {
	*baseHttpService
}

func (s *oldWebService) Name() string {
	return "OldWebService"
}

func (s *oldWebService) Init() error {
	s.registerRoute(routers.RegisterRoute)
	return nil
}

func (s *oldWebService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.ServerSetting.HttpIp, conf.ServerSetting.HttpPort))
}

func newOldWebService() Service {
	addr := conf.ServerSetting.HttpIp + ":" + conf.ServerSetting.HttpPort
	server := httpServerFrom(addr, func() *httpServer {
		engine := newWebEngine()
		return &httpServer{
			e: engine,
			server: &http.Server{
				Addr:           addr,
				Handler:        engine,
				ReadTimeout:    conf.ServerSetting.ReadTimeout,
				WriteTimeout:   conf.ServerSetting.WriteTimeout,
				MaxHeaderBytes: 1 << 20,
			},
		}
	})
	return &oldWebService{
		baseHttpService: &baseHttpService{
			server: server,
		},
	}
}
