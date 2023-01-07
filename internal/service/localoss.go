// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"net/http"

	"github.com/Masterminds/semver/v3"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/servants"
)

var (
	_ Service = (*localossService)(nil)
)

type localossService struct {
	*baseHttpService
}

func (s *localossService) Name() string {
	return "LocalossService"
}

func (s *localossService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *localossService) OnInit() error {
	s.registerRoute(s, servants.RegisterLocalossServants)
	return nil
}

func (s *localossService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.LocalossServerSetting.HttpIp, conf.LocalossServerSetting.HttpPort))
}

func newLocalossEngine() *gin.Engine {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	return e
}

func newLocalossService() Service {
	addr := conf.LocalossServerSetting.HttpIp + ":" + conf.LocalossServerSetting.HttpPort
	server := httpServers.from(addr, func() *httpServer {
		engine := newLocalossEngine()
		return &httpServer{
			baseServer: newBaseServe(),
			e:          engine,
			server: &http.Server{
				Addr:           addr,
				Handler:        engine,
				ReadTimeout:    conf.LocalossServerSetting.GetReadTimeout(),
				WriteTimeout:   conf.LocalossServerSetting.GetWriteTimeout(),
				MaxHeaderBytes: 1 << 20,
			},
		}
	})
	return &localossService{
		baseHttpService: &baseHttpService{
			server: server,
		},
	}
}
