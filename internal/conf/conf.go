// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"log"
	"time"

	"github.com/alimy/tryst/cfg"
)

var (
	loggerSetting            *loggerConf
	loggerFileSetting        *loggerFileConf
	loggerZincSetting        *loggerZincConf
	loggerMeiliSetting       *loggerMeiliConf
	loggerOpenObserveSetting *loggerOpenObserveConf
	loggerOtlpSetting        *loggerOtlponf
	sentrySetting            *sentryConf
	redisSetting             *redisConf

	PyroscopeSetting        *pyroscopeConf
	DatabaseSetting         *databaseConf
	MysqlSetting            *mysqlConf
	PostgresSetting         *postgresConf
	Sqlite3Setting          *sqlite3Conf
	PprofServerSetting      *httpServerConf
	MetricsServerSetting    *httpServerConf
	WebServerSetting        *httpServerConf
	AdminServerSetting      *httpServerConf
	SpaceXServerSetting     *httpServerConf
	BotServerSetting        *httpServerConf
	LocalossServerSetting   *httpServerConf
	FrontendWebSetting      *httpServerConf
	DocsServerSetting       *httpServerConf
	MobileServerSetting     *grpcServerConf
	AppSetting              *appConf
	CacheSetting            *cacheConf
	EventManagerSetting     *eventManagerConf
	MetricManagerSetting    *metricManagerConf
	JobManagerSetting       *jobManagerConf
	CacheIndexSetting       *cacheIndexConf
	SimpleCacheIndexSetting *simpleCacheIndexConf
	BigCacheIndexSetting    *bigCacheIndexConf
	RedisCacheIndexSetting  *redisCacheIndexConf
	SmsJuheSetting          *smsJuheConf
	AlipaySetting           *alipayConf
	TweetSearchSetting      *tweetSearchConf
	ZincSetting             *zincConf
	MeiliSetting            *meiliConf
	ObjectStorage           *objectStorageConf
	AliOSSSetting           *aliOSSConf
	COSSetting              *cosConf
	HuaweiOBSSetting        *huaweiOBSConf
	MinIOSetting            *minioConf
	S3Setting               *s3Conf
	LocalOSSSetting         *localossConf
	JWTSetting              *jwtConf
	WebProfileSetting       *WebProfileConf
)

func setupSetting(suite []string, noDefault bool) error {
	vp, err := newViper()
	if err != nil {
		return err
	}

	// initialize features configure
	ss, kv := featuresInfoFrom(vp, "Features")
	cfg.Initial(ss, kv)
	if len(suite) > 0 {
		cfg.Use(suite, noDefault)
	}

	objects := map[string]any{
		"App":               &AppSetting,
		"Cache":             &CacheSetting,
		"EventManager":      &EventManagerSetting,
		"MetricManager":     &MetricManagerSetting,
		"JobManager":        &JobManagerSetting,
		"PprofServer":       &PprofServerSetting,
		"MetricsServer":     &MetricsServerSetting,
		"WebServer":         &WebServerSetting,
		"AdminServer":       &AdminServerSetting,
		"SpaceXServer":      &SpaceXServerSetting,
		"BotServer":         &BotServerSetting,
		"LocalossServer":    &LocalossServerSetting,
		"FrontendWebServer": &FrontendWebSetting,
		"DocsServer":        &DocsServerSetting,
		"MobileServer":      &MobileServerSetting,
		"CacheIndex":        &CacheIndexSetting,
		"SimpleCacheIndex":  &SimpleCacheIndexSetting,
		"BigCacheIndex":     &BigCacheIndexSetting,
		"RedisCacheIndex":   &RedisCacheIndexSetting,
		"Alipay":            &AlipaySetting,
		"SmsJuhe":           &SmsJuheSetting,
		"Pyroscope":         &PyroscopeSetting,
		"Sentry":            &sentrySetting,
		"Logger":            &loggerSetting,
		"LoggerFile":        &loggerFileSetting,
		"LoggerZinc":        &loggerZincSetting,
		"LoggerMeili":       &loggerMeiliSetting,
		"LoggerOpenObserve": &loggerOpenObserveSetting,
		"LoggerOtlp":        &loggerOtlpSetting,
		"Database":          &DatabaseSetting,
		"MySQL":             &MysqlSetting,
		"Postgres":          &PostgresSetting,
		"Sqlite3":           &Sqlite3Setting,
		"TweetSearch":       &TweetSearchSetting,
		"Zinc":              &ZincSetting,
		"Meili":             &MeiliSetting,
		"Redis":             &redisSetting,
		"JWT":               &JWTSetting,
		"ObjectStorage":     &ObjectStorage,
		"AliOSS":            &AliOSSSetting,
		"COS":               &COSSetting,
		"HuaweiOBS":         &HuaweiOBSSetting,
		"MinIO":             &MinIOSetting,
		"LocalOSS":          &LocalOSSSetting,
		"S3":                &S3Setting,
		"WebProfile":        &WebProfileSetting,
	}
	for k, v := range objects {
		err := vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}

	CacheSetting.CientSideCacheExpire *= time.Second
	EventManagerSetting.MaxIdleTime *= time.Second
	MetricManagerSetting.MaxIdleTime *= time.Second
	JWTSetting.Expire *= time.Second
	SimpleCacheIndexSetting.CheckTickDuration *= time.Second
	SimpleCacheIndexSetting.ExpireTickDuration *= time.Second
	BigCacheIndexSetting.ExpireInSecond *= time.Second
	RedisCacheIndexSetting.ExpireInSecond *= time.Second
	redisSetting.ConnWriteTimeout *= time.Second

	return nil
}

func Initial(suite []string, noDefault bool) {
	err := setupSetting(suite, noDefault)
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	setupLogger()
	initSentry()
}

func GetOssDomain() string {
	uri := "https://"
	if cfg.If("AliOSS") {
		return uri + AliOSSSetting.Domain + "/"
	} else if cfg.If("COS") {
		return uri + COSSetting.Domain + "/"
	} else if cfg.If("HuaweiOBS") {
		return uri + HuaweiOBSSetting.Domain + "/"
	} else if cfg.If("MinIO") {
		if !MinIOSetting.Secure {
			uri = "http://"
		}
		return uri + MinIOSetting.Domain + "/" + MinIOSetting.Bucket + "/"
	} else if cfg.If("S3") {
		if !S3Setting.Secure {
			uri = "http://"
		}
		// TODO: will not work well need test in real world
		return uri + S3Setting.Domain + "/" + S3Setting.Bucket + "/"
	} else if cfg.If("LocalOSS") {
		if !LocalOSSSetting.Secure {
			uri = "http://"
		}
		return uri + LocalOSSSetting.Domain + "/oss/" + LocalOSSSetting.Bucket + "/"
	}
	return uri + AliOSSSetting.Domain + "/"
}

func RunMode() string {
	return AppSetting.RunMode
}

func UseSentryGin() bool {
	return cfg.If("Sentry") && sentrySetting.AttachGin
}
