// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/fatih/color"
	"github.com/rocboss/paopao-ce/internal/servants"
)

var (
	_ Service = (*tripletService)(nil)
)

// tripletService just a demo service for Connect
type tripletService struct {
	*baseConnectService
}

func (s *tripletService) Name() string {
	return "TripletService"
}

func (s *tripletService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *tripletService) OnInit() error {
	s.registerServer(s, servants.RegisterTripletServants)
	return nil
}

func (s *tripletService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://localhost:8080"))
}

func newTripletService() Service {
	server := connectServers.from(":8080", func() *connectServer {
		// TODO: do something for conenct server initialize
		return defaultConnectServer(":8080")
	})
	return &tripletService{
		baseConnectService: &baseConnectService{
			server: server,
		},
	}
}
