package conf

import (
	"log"
	"sync"
	"time"

	"github.com/rocboss/paopao-ce/pkg/cfg"
)

var (
	loggerSetting      *LoggerSettingS
	loggerFileSetting  *LoggerFileSettingS
	loggerZincSetting  *LoggerZincSettingS
	loggerMeiliSetting *LoggerMeiliSettingS
	redisSetting       *RedisSettingS

	DatabaseSetting         *DatabaseSetingS
	MysqlSetting            *MySQLSettingS
	PostgresSetting         *PostgresSettingS
	Sqlite3Setting          *Sqlite3SettingS
	ServerSetting           *ServerSettingS
	AppSetting              *AppSettingS
	CacheIndexSetting       *CacheIndexSettingS
	SimpleCacheIndexSetting *SimpleCacheIndexSettingS
	BigCacheIndexSetting    *BigCacheIndexSettingS
	SmsJuheSetting          *SmsJuheSettings
	AlipaySetting           *AlipaySettingS
	TweetSearchSetting      *TweetSearchS
	ZincSetting             *ZincSettingS
	MeiliSetting            *MeiliSettingS
	ObjectStorage           *ObjectStorageS
	AliOSSSetting           *AliOSSSettingS
	COSSetting              *COSSettingS
	HuaweiOBSSetting        *HuaweiOBSSettingS
	MinIOSetting            *MinIOSettingS
	S3Setting               *S3SettingS
	LocalOSSSetting         *LocalOSSSettingS
	JWTSetting              *JWTSettingS
	Mutex                   *sync.Mutex
)

func setupSetting(suite []string, noDefault bool) error {
	setting, err := NewSetting()
	if err != nil {
		return err
	}

	// initialize features configure
	ss, kv := setting.featuresInfoFrom("Features")
	cfg.Initialize(ss, kv)
	if len(suite) > 0 {
		cfg.Use(suite, noDefault)
	}

	objects := map[string]any{
		"App":              &AppSetting,
		"Server":           &ServerSetting,
		"CacheIndex":       &CacheIndexSetting,
		"SimpleCacheIndex": &SimpleCacheIndexSetting,
		"BigCacheIndex":    &BigCacheIndexSetting,
		"Alipay":           &AlipaySetting,
		"SmsJuhe":          &SmsJuheSetting,
		"Logger":           &loggerSetting,
		"LoggerFile":       &loggerFileSetting,
		"LoggerZinc":       &loggerZincSetting,
		"LoggerMeili":      &loggerMeiliSetting,
		"Database":         &DatabaseSetting,
		"MySQL":            &MysqlSetting,
		"Postgres":         &PostgresSetting,
		"Sqlite3":          &Sqlite3Setting,
		"TweetSearch":      &TweetSearchSetting,
		"Zinc":             &ZincSetting,
		"Meili":            &MeiliSetting,
		"Redis":            &redisSetting,
		"JWT":              &JWTSetting,
		"ObjectStorage":    &ObjectStorage,
		"AliOSS":           &AliOSSSetting,
		"COS":              &COSSetting,
		"HuaweiOBS":        &HuaweiOBSSetting,
		"MinIO":            &MinIOSetting,
		"LocalOSS":         &LocalOSSSetting,
		"S3":               &S3Setting,
	}
	if err = setting.Unmarshal(objects); err != nil {
		return err
	}

	JWTSetting.Expire *= time.Second
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second
	SimpleCacheIndexSetting.CheckTickDuration *= time.Second
	SimpleCacheIndexSetting.ExpireTickDuration *= time.Second
	BigCacheIndexSetting.ExpireInSecond *= time.Second

	Mutex = &sync.Mutex{}
	return nil
}

func Initialize(suite []string, noDefault bool) {
	err := setupSetting(suite, noDefault)
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	setupLogger()
	setupDBEngine()
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
