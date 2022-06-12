package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/routers"
	"github.com/rocboss/paopao-ce/pkg/util"
)

var (
	version, commitID, buildDate string

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

func main() {
	gin.SetMode(conf.ServerSetting.RunMode)

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           conf.ServerSetting.HttpIp + ":" + conf.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    conf.ServerSetting.ReadTimeout,
		WriteTimeout:   conf.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	util.PrintHelloBanner(fmt.Sprintf("paopao %s (build:%s %s)", version, commitID, buildDate))
	fmt.Fprintf(color.Output, "PaoPao service listen on %s\n",
		color.GreenString(fmt.Sprintf("http://%s:%s", conf.ServerSetting.HttpIp, conf.ServerSetting.HttpPort)),
	)
	s.ListenAndServe()
}
