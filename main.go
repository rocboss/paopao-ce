// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/service"
	"github.com/rocboss/paopao-ce/pkg/debug"
	"github.com/rocboss/paopao-ce/pkg/util"
)

var (
	noDefaultFeatures bool
	features          suites
)

type suites []string

func (s *suites) String() string {
	return strings.Join(*s, ",")
}

func (s *suites) Set(value string) error {
	for _, item := range strings.Split(value, ",") {
		*s = append(*s, strings.TrimSpace(item))
	}
	return nil
}

func init() {
	flagParse()

	conf.Initialize(features, noDefaultFeatures)
	internal.Initialize()
}

func flagParse() {
	flag.BoolVar(&noDefaultFeatures, "no-default-features", false, "whether use default features")
	flag.Var(&features, "features", "use special features")
	flag.Parse()
}

func runService(ss []service.Service) {
	gin.SetMode(conf.RunMode())

	fmt.Fprintf(color.Output, "\nstarting run service...\n\n")
	for _, s := range ss {
		go func(s service.Service) {
			fmt.Fprintf(color.Output, "%s[start] - %s\n", s.Name(), s)
			if err := s.Start(); err != nil {
				log.Fatalf("%s[start] - occurs on error: %s\n", s.Name(), err)
			}
		}(s)
	}

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Fprintf(color.Output, "\nshutting down server...\n\n")

	for _, s := range ss {
		if err := s.Stop(); err != nil {
			fmt.Fprintf(color.Output, "%s[stop] - occurs on error: %s\n", s.Name(), err)
		}
		fmt.Fprintf(color.Output, "%s[stop] - finish...\n", s.Name())
	}
}

func main() {
	util.PrintHelloBanner(debug.VersionInfo())

	if ss := service.InitService(); len(ss) > 0 {
		runService(ss)
	} else {
		fmt.Fprintf(color.Output, "no service need start so just exit")
	}
}
