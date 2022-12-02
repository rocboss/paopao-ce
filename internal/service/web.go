// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/servants"
)

type webService struct {
	baseService

	server *http.Server
}

func (s *webService) Name() string {
	return "WebService"
}

func (s *webService) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("run app failed: %s", err)
	}
}

func (s *webService) Stop() {
	s.server.Shutdown(context.Background())
}

func (s *webService) String() string {
	return fmt.Sprintf("PaoPao service listen on %s\n", color.GreenString("http://%s:%s", conf.ServerSetting.HttpIp, conf.ServerSetting.HttpPort))
}

func newWebEngine() *gin.Engine {
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

func NewWebService() Service {
	e := newWebEngine()
	servants.RegisterWebServants(e)
	return &webService{
		server: &http.Server{
			Addr:           conf.ServerSetting.HttpIp + ":" + conf.ServerSetting.HttpPort,
			Handler:        e,
			ReadTimeout:    conf.ServerSetting.ReadTimeout,
			WriteTimeout:   conf.ServerSetting.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}
}
