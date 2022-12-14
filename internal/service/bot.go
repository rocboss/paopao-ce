// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/servants"
)

type botService struct {
	*baseHttpService
}

func (s *botService) Name() string {
	return "BotService"
}

func (s *botService) Init() error {
	s.registerRoute(servants.RegisterBotServants)
	return nil
}

func (s *botService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.BotServerSetting.HttpIp, conf.BotServerSetting.HttpPort))
}

func newBotEngine() *gin.Engine {
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

func newBotService() Service {
	addr := conf.BotServerSetting.HttpIp + ":" + conf.BotServerSetting.HttpPort
	server := httpServerFrom(addr, func() *httpServer {
		engine := newBotEngine()
		return &httpServer{
			e: engine,
			server: &http.Server{
				Addr:           addr,
				Handler:        engine,
				ReadTimeout:    conf.BotServerSetting.ReadTimeout,
				WriteTimeout:   conf.BotServerSetting.WriteTimeout,
				MaxHeaderBytes: 1 << 20,
			},
		}
	})
	return &botService{
		baseHttpService: &baseHttpService{
			server: server,
		},
	}
}
