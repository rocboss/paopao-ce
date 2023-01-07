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
	_ Service = (*docsService)(nil)
)

type docsService struct {
	*baseHttpService
}

func (s *docsService) Name() string {
	return "DocsService"
}

func (s *docsService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *docsService) OnInit() error {
	s.registerRoute(s, servants.RegisterDocsServants)
	return nil
}

func (s *docsService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.DocsServerSetting.HttpIp, conf.DocsServerSetting.HttpPort))
}

func newDocsService() Service {
	addr := conf.DocsServerSetting.HttpIp + ":" + conf.DocsServerSetting.HttpPort
	server := httpServers.from(addr, func() *httpServer {
		engine := newWebEngine()
		return &httpServer{
			baseServer: newBaseServe(),
			e:          engine,
			server: &http.Server{
				Addr:           addr,
				Handler:        engine,
				ReadTimeout:    conf.DocsServerSetting.GetReadTimeout(),
				WriteTimeout:   conf.DocsServerSetting.GetWriteTimeout(),
				MaxHeaderBytes: 1 << 20,
			},
		}
	})
	return &docsService{
		baseHttpService: &baseHttpService{
			server: server,
		},
	}
}
