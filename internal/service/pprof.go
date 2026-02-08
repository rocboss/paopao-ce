// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"net/http"

	"github.com/Masterminds/semver/v3"
	"github.com/fatih/color"
	"github.com/rocboss/paopao-ce/internal/conf"
)

var _ Service = (*pprofService)(nil)

type pprofService struct {
	*baseHttpService
}

func (s *pprofService) Name() string {
	return "PprofService"
}

func (s *pprofService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *pprofService) OnInit() error {
	s.registerRoute(s, nil)
	return nil
}

func (s *pprofService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.PprofServerSetting.HttpIp, conf.PprofServerSetting.HttpPort))
}

func newPprofService() Service {
	addr := conf.PprofServerSetting.HttpIp + ":" + conf.PprofServerSetting.HttpPort
	// notice this step just to register pprof server to start. don't share server with pprof.
	server := httpServers.from(addr, func() *httpServer {
		return &httpServer{
			baseServer: newBaseServe(),
			server: &http.Server{
				Addr:    addr,
				Handler: http.DefaultServeMux,
			},
		}
	})
	return &pprofService{
		baseHttpService: &baseHttpService{
			server: server,
		},
	}
}
