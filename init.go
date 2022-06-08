package main

import (
	"flag"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/routers/api"
	"github.com/rocboss/paopao-ce/internal/service"
	"github.com/rocboss/paopao-ce/pkg/logger"
	"github.com/rocboss/paopao-ce/pkg/setting"
	"github.com/rocboss/paopao-ce/pkg/zinc"
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

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	client := zinc.NewClient(global.ZincSetting)
	service.Initialize(global.DBEngine, client)
	api.Initialize()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	global.Features = setting.FeaturesFrom("Features")
	if len(features) > 0 {
		if err = global.Features.Use(features, noDefaultFeatures); err != nil {
			return err
		}
	}

	objects := map[string]interface{}{
		"App":        &global.AppSetting,
		"Server":     &global.ServerSetting,
		"Alipay":     &global.AlipaySetting,
		"SmsJuhe":    &global.SmsJuheSetting,
		"LoggerFile": &global.LoggerFileSetting,
		"LoggerZinc": &global.LoggerZincSetting,
		"MySQL":      &global.MySQLSetting,
		"Zinc":       &global.ZincSetting,
		"Redis":      &global.RedisSetting,
		"JWT":        &global.JWTSetting,
		"AliOSS":     &global.AliOSSSetting,
		"MinIO":      &global.MinIOSetting,
		"S3":         &global.S3Setting,
	}
	if err = setting.Unmarshal(objects); err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.Mutex = &sync.Mutex{}
	return nil
}

func flagParse() {
	flag.BoolVar(&noDefaultFeatures, "no-default-features", false, "whether use default features")
	flag.Var(&features, "features", "use special features")
	flag.Parse()
}

func setupLogger() error {
	logger, err := logger.New()
	if err != nil {
		return err
	}
	global.Logger = logger

	return nil
}

// setupDBEngine 暂时只支持MySQL
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.MySQLSetting)
	if err != nil {
		return err
	}

	global.Redis = redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Host,
		Password: global.RedisSetting.Password,
		DB:       global.RedisSetting.DB,
	})

	return nil
}
