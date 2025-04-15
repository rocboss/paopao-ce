// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package serve

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alimy/tryst/cfg"
	"github.com/fatih/color"
	"github.com/getsentry/sentry-go"
	"github.com/rocboss/paopao-ce/cmd"
	"github.com/rocboss/paopao-ce/internal"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/service"
	"github.com/rocboss/paopao-ce/pkg/debug"
	"github.com/rocboss/paopao-ce/pkg/utils"
	"github.com/rocboss/paopao-ce/pkg/version"
	"github.com/sourcegraph/conc"
	"github.com/spf13/cobra"
	"go.uber.org/automaxprocs/maxprocs"
)

var (
	noDefaultFeatures bool
	features          []string
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start paopao-ce server",
		Long:  "start paopao-ce server",
		Run:   serveRun,
	}

	serveCmd.Flags().BoolVar(&noDefaultFeatures, "no-default-features", false, "whether not use default features")
	serveCmd.Flags().StringSliceVarP(&features, "features", "f", []string{}, "use special features")

	cmd.Register(serveCmd)
}

func deferFn() {
	if cfg.If("Sentry") {
		// Flush buffered events before the program terminates.
		sentry.Flush(2 * time.Second)
	}
	conf.CloseDB()
}

func serveRun(_cmd *cobra.Command, _args []string) {
	utils.PrintHelloBanner(version.VersionInfo())

	// set maxprocs automatic
	maxprocs.Set(maxprocs.Logger(log.Printf))

	// initial configure
	conf.Initial(features, noDefaultFeatures)
	if cfg.If("loggerOtlp") {
		shutdownFn, _ := conf.InitTelemetry()
		defer shutdownFn()
	}
	internal.Initial()
	ss := service.MustInitService()
	if len(ss) < 1 {
		fmt.Fprintln(color.Output, "no service need start so just exit")
		return
	}

	// do defer function
	defer deferFn()

	// start pyroscope if need
	debug.StartPyroscope()

	// start services
	wg := conc.NewWaitGroup()
	fmt.Fprintf(color.Output, "\nstarting run service...\n\n")
	service.Start(wg)

	// graceful stop services
	wg.Go(func() {
		quit := make(chan os.Signal, 1)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Fprintf(color.Output, "\nshutting down server...\n\n")
		service.Stop()
	})
	wg.Wait()
}
