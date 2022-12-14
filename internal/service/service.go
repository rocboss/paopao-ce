// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"log"

	"github.com/alimy/cfg"
	"github.com/rocboss/paopao-ce/pkg/types"
)

type Service interface {
	Name() string
	Init() error
	Start() error
	Stop() error
}

type baseService types.Empty

func (baseService) Name() string {
	return ""
}

func (baseService) String() string {
	return ""
}

func InitService() []Service {
	ss := newService()
	for _, s := range ss {
		if err := s.Init(); err != nil {
			log.Fatalf("initial %s service error: %s", s.Name(), err)
		}
	}
	return ss
}

func newService() (ss []Service) {
	ss = append(ss, newWebService())

	// add oldWebService if not depredcated OldWebService
	cfg.Not("Deprecated:OldWeb", func() {
		ss = append(ss, newOldWebService())
	})

	cfg.In(cfg.Actions{
		"Admin": func() {
			ss = append(ss, newAdminService())
		},
		"SpaceX": func() {
			ss = append(ss, newSpaceXService())
		},
		"Bot": func() {
			ss = append(ss, newBotService())
		},
		"LocalOSS": func() {
			ss = append(ss, newLocalossService())
		},
	})

	return
}
