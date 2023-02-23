// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"net/http"

	"github.com/Masterminds/semver/v3"
	"github.com/fatih/color"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/servants"
)

var (
	_ Service = (*frontendWebService)(nil)
)

type frontendWebService struct {
	*baseHttpService
}

func (s *frontendWebService) Name() string {
	return "FrontendWebService"
}

func (s *frontendWebService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *frontendWebService) OnInit() error {
	s.registerRoute(s, servants.RegisterFrontendWebServants)
	return nil
}

func (s *frontendWebService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.FrontendWebSetting.HttpIp, conf.FrontendWebSetting.HttpPort))
}

func newFrontendWebServiceService() Service {
	addr := conf.FrontendWebSetting.HttpIp + ":" + conf.FrontendWebSetting.HttpPort
	server := httpServers.from(addr, func() *httpServer {
		engine := newWebEngine()
		return &httpServer{
			baseServer: newBaseServe(),
			e:          engine,
			server: &http.Server{
				Addr:           addr,
				Handler:        engine,
				ReadTimeout:    conf.FrontendWebSetting.GetReadTimeout(),
				WriteTimeout:   conf.FrontendWebSetting.GetWriteTimeout(),
				MaxHeaderBytes: 1 << 20,
			},
		}
	})
	return &frontendWebService{
		baseHttpService: &baseHttpService{
			server: server,
		},
	}
}
