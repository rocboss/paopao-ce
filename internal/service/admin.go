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
	_ Service = (*adminService)(nil)
)

type adminService struct {
	*baseHttpService
}

func (s *adminService) Name() string {
	return "AdminService"
}

func (s *adminService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *adminService) OnInit() error {
	s.registerRoute(s, servants.RegisterAdminServants)
	return nil
}

func (s *adminService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.AdminServerSetting.HttpIp, conf.AdminServerSetting.HttpPort))
}

func newAdminEngine() *gin.Engine {
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

func newAdminService() Service {
	addr := conf.AdminServerSetting.HttpIp + ":" + conf.AdminServerSetting.HttpPort
	server := httpServers.from(addr, func() *httpServer {
		engine := newAdminEngine()
		return &httpServer{
			baseServer: newBaseServe(),
			e:          engine,
			server: &http.Server{
				Addr:           addr,
				Handler:        engine,
				ReadTimeout:    conf.AdminServerSetting.GetReadTimeout(),
				WriteTimeout:   conf.AdminServerSetting.GetWriteTimeout(),
				MaxHeaderBytes: 1 << 20,
			},
		}
	})
	return &adminService{
		baseHttpService: &baseHttpService{
			server: server,
		},
	}
}
