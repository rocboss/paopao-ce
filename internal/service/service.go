// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"log"

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

func InitService() (ss []Service) {
	ss = append(ss, newWebService(), newOldWebService())

	for _, s := range ss {
		if err := s.Init(); err != nil {
			log.Fatalf("initial %s service error: %s", s.Name(), err)
		}
	}
	return
}
