package app

import (
	"net/http"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/servants"
)

type Service interface {
	Start()
	Stop()
	Info() string
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
