// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
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
	flag.BoolVar(&noDefaultFeatures, "no-default-features", false, "whether not use default features")
	flag.Var(&features, "features", "use special features")
	flag.Parse()
}

func runService(wg *sync.WaitGroup, ss []service.Service) {
	gin.SetMode(conf.RunMode())

	fmt.Fprintf(color.Output, "\nstarting run service...\n\n")
	l := service.MaxSidSize(ss)
	for _, s := range ss {
		go func(s service.Service) {
			fmt.Fprintf(color.Output, "%s [start] - %s", util.SidStr(s.Name(), s.Version(), l), s)
			if err := s.OnStart(); err != nil {
				fmt.Fprintf(color.Output, "%s [start] - occurs on error: %s\n", util.SidStr(s.Name(), s.Version(), l), err)
			}
			wg.Done()
		}(s)
	}
}

func runManage(wg *sync.WaitGroup, ss []service.Service) {
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Fprintf(color.Output, "\nshutting down server...\n\n")
	l := service.MaxSidSize(ss)
	for _, s := range ss {
		if err := s.OnStop(); err != nil {
			fmt.Fprintf(color.Output, "%s [stop]  - occurs on error: %s\n", util.SidStr(s.Name(), s.Version(), l), err)
		}
		fmt.Fprintf(color.Output, "%s [stop]  - finish...\n", util.SidStr(s.Name(), s.Version(), l))
	}
	wg.Done()
}

func main() {
	util.PrintHelloBanner(debug.VersionInfo())

	if ss := service.InitService(); len(ss) > 0 {
		wg := &sync.WaitGroup{}
		wg.Add(len(ss) + 1)

		runService(wg, ss)
		go runManage(wg, ss)

		wg.Wait()
	} else {
		fmt.Fprintln(color.Output, "no service need start so just exit")
	}
}
