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

func main() {
	util.PrintHelloBanner(debug.VersionInfo())
	ss := service.MustInitService()
	if len(ss) < 1 {
		fmt.Fprintln(color.Output, "no service need start so just exit")
		return
	}
	wg := &sync.WaitGroup{}
	// start services
	fmt.Fprintf(color.Output, "\nstarting run service...\n\n")
	service.Start(wg)
	// graceful stop services
	wg.Add(1)
	go func() {
		quit := make(chan os.Signal, 1)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Fprintf(color.Output, "\nshutting down server...\n\n")
		service.Stop()
		wg.Done()
	}()
	wg.Wait()
}
