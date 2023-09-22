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
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/internal/metrics/prometheus"
)

var (
	_ Service = (*metricsService)(nil)
)

type metricsService struct {
	*baseHttpService
}

func (s *metricsService) Name() string {
	return "MetricsService"
}

func (s *metricsService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *metricsService) OnInit() error {
	s.registerRoute(s, nil)
	return nil
}

func (s *metricsService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.MetricsServerSetting.HttpIp, conf.MetricsServerSetting.HttpPort))
}

func newMetricsService() Service {
	addr := conf.MetricsServerSetting.HttpIp + ":" + conf.MetricsServerSetting.HttpPort
	// notice this step just to register pprof server to start. don't share server with pprof.
	server := httpServers.from(addr, func() *httpServer {
		ds, wc := dao.DataService(), cache.NewWebCache()
		mux := http.NewServeMux()
		mux.Handle("/metrics", prometheus.NewHandler(ds, wc))
		return &httpServer{
			baseServer: newBaseServe(),
			server: &http.Server{
				Addr:    addr,
				Handler: mux,
			},
		}
	})
	return &metricsService{
		baseHttpService: &baseHttpService{
			server: server,
		},
	}
}
