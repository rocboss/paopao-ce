// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build pyroscope
// +build pyroscope

package debug

import (
	"os"

	"github.com/alimy/tryst/cfg"
	pyroscope "github.com/grafana/pyroscope-go"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

func StartPyroscope() {
	if !cfg.If("Pyroscope") {
		logrus.Infoln("skip Pyroscope because not add Pyroscope feature in config.yaml")
		return
	}
	s := conf.PyroscopeSetting
	c := pyroscope.Config{
		ApplicationName: s.AppName,
		ServerAddress:   s.Endpoint,
		AuthToken:       os.Getenv("PYROSCOPE_AUTH_TOKEN"),
		Logger:          s.GetLogger(),
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
		},
	}
	if len(c.AuthToken) == 0 {
		c.AuthToken = s.AuthToken
	}
	pyroscope.Start(c)
}
