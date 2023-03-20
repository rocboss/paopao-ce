// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"log"

	"github.com/Masterminds/semver/v3"
	"github.com/alimy/cfg"
	"github.com/rocboss/paopao-ce/pkg/types"
)

type Service interface {
	Name() string
	Version() *semver.Version
	OnInit() error
	OnStart() error
	OnStop() error
}

type baseService types.Empty

func (baseService) Name() string {
	return ""
}

func (baseService) Version() *semver.Version {
	return semver.MustParse("v0.0.1")
}

func (baseService) String() string {
	return ""
}

// MustInitService Initial service
func MustInitService() []Service {
	ss := newService()
	for _, s := range ss {
		if err := s.OnInit(); err != nil {
			log.Fatalf("initial %s service error: %s", s.Name(), err)
		}
	}
	return ss
}

func newService() (ss []Service) {
	// add all service if declared in features on config.yaml
	cfg.In(cfg.Actions{
		"Web": func() {
			ss = append(ss, newWebService())
		},
		"Admin": func() {
			ss = append(ss, newAdminService())
		},
		"SpaceX": func() {
			ss = append(ss, newSpaceXService())
		},
		"Bot": func() {
			ss = append(ss, newBotService())
		},
		"NativeOBS": func() {
			ss = append(ss, newLocalossService())
		},
		"Mobile": func() {
			ss = append(ss, newMobileService())
		},
		"Frontend:Web": func() {
			ss = append(ss, newFrontendWebServiceService())
		},
		"Docs": func() {
			ss = append(ss, newDocsService())
		},
	})
	return
}
