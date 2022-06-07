package main

import (
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/service"
	"github.com/rocboss/paopao-ce/pkg/logger"
	"github.com/rocboss/paopao-ce/pkg/setting"
	"github.com/rocboss/paopao-ce/pkg/zinc"
)

func init() {
	const RESTARTS = 5
	const SECONDS = 10
	var err error

	// Settings and logger have to be set up synchronously
	// because they modify global state
	// before we can setup db and search engine.
	if err = setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	if err = setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	// DB and search engine are set up concurrently with restarts,
	// because they depend on external services.
	// Although the <- operations are blocking,
	// these functions are all executing in the background when other <- operations are run
	// thus they are in fact concurrent. (Assuming <- has negligible cost, which it has).
	dbOk := setupAsync(setupDBEngine, RESTARTS, SECONDS)
	searchOk := setupAsync(setupSearchEngine, RESTARTS, SECONDS)

	if err = <-dbOk; err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	if err = <-searchOk; err != nil {
		log.Fatalf("init.setupSearchEngine err: %v", err)
	}
}

// setupAsync calls a setup function asynchronously, with restarts,
// returning a channel as an error handle.
func setupAsync(setupFunc func() error, restarts, seconds int) <-chan error {
	errChan := make(chan error)

	// This function re-executes the setupFunc until it succeeds (err == nil)
	// or the function was executed for restarts times.
	// In case all restarts attemtps fail, the last error is sent back in the channel.
	go func() {
		var err error
		for i := 0; i < restarts; i++ {
			if err = setupFunc(); err == nil {
				errChan <- nil
				return
			} else {
				log.Printf("%v encountered. Restarting...", err)
			}

			time.Sleep(time.Duration(seconds) * time.Second)
		}
		errChan <- err
	}()

	return errChan
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	global.Features = setting.FeaturesFrom("Features")

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

// setupSearchEngine only returns nil for now, however,
// it could be modified to return an error in the future
func setupSearchEngine() error {
	client := zinc.NewClient(global.ZincSetting)
	service.Initialize(global.DBEngine, client)
	return nil
}
