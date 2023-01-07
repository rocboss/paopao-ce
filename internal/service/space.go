// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"net/http"

	"github.com/Masterminds/semver/v3"
	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/servants"
)

var (
	_ Service = (*spaceXService)(nil)
)

type spaceXService struct {
	*baseHttpService
}

func (s *spaceXService) Name() string {
	return "SpaceXService"
}

func (s *spaceXService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *spaceXService) OnInit() error {
	s.registerRoute(s, servants.RegisterSpaceXServants)
	return nil
}

func (s *spaceXService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.SpaceXServerSetting.HttpIp, conf.SpaceXServerSetting.HttpPort))
}

func newSpaceXEngine() *gin.Engine {
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	// 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	e.Use(cors.New(corsConfig))

	// 默认404
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})

	// 默认405
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "Method Not Allowed",
		})
	})

	return e
}

func newSpaceXService() Service {
	addr := conf.SpaceXServerSetting.HttpIp + ":" + conf.SpaceXServerSetting.HttpPort
	server := httpServers.from(addr, func() *httpServer {
		engine := newSpaceXEngine()
		return &httpServer{
			baseServer: newBaseServe(),
			e:          engine,
			server: &http.Server{
				Addr:           addr,
				Handler:        engine,
				ReadTimeout:    conf.SpaceXServerSetting.GetReadTimeout(),
				WriteTimeout:   conf.SpaceXServerSetting.GetWriteTimeout(),
				MaxHeaderBytes: 1 << 20,
			},
		}
	})
	return &spaceXService{
		baseHttpService: &baseHttpService{
			server: server,
		},
	}
}
