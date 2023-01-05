// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"log"
	"net"

	"github.com/Masterminds/semver/v3"
	"github.com/fatih/color"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/servants"
	"google.golang.org/grpc"
)

var (
	_ Service = (*mobileService)(nil)
)

type mobileService struct {
	*baseGRPCService
}

func (s *mobileService) Name() string {
	return "MobileService"
}

func (s *mobileService) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *mobileService) OnInit() error {
	s.registerServer(s, servants.RegisterMobileServants)
	return nil
}

func (s *mobileService) String() string {
	return fmt.Sprintf("listen on %s\n", color.GreenString("http://%s:%s", conf.MobileServerSetting.Host, conf.MobileServerSetting.Port))
}

func newMobileService() Service {
	addr := conf.MobileServerSetting.Host + ":" + conf.MobileServerSetting.Port
	server := grpcServers.from(addr, func() *grpcServer {
		l, err := net.Listen("tcp", addr)
		if err != nil {
			// TODO: optimize error process
			log.Fatal(fmt.Errorf("failed to listen: %w", err))
		}

		// TODO: init grpc.Server
		s := grpc.NewServer()

		return &grpcServer{
			baseServer: newBaseServe(),
			listener:   l,
			server:     s,
		}
	})
	return &mobileService{
		baseGRPCService: &baseGRPCService{
			server: server,
		},
	}
}
