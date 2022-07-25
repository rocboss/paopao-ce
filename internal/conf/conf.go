package conf

import (
	"log"
	"sync"
	"time"
)

var (
	loggerSetting      *LoggerSettingS
	loggerFileSetting  *LoggerFileSettingS
	loggerZincSetting  *LoggerZincSettingS
	loggerMeiliSetting *LoggerMeiliSettingS
	redisSetting       *RedisSettingS
	features           *FeaturesSettingS

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

	features = setting.FeaturesFrom("Features")
	if len(suite) > 0 {
		if err = features.Use(suite, noDefault); err != nil {
			return err
		}
	}

	objects := map[string]interface{}{
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

// Cfg get value by key if exist
func Cfg(key string) (string, bool) {
	return features.Cfg(key)
}

// CfgIf check expression is true. if expression just have a string like
// `Sms` is mean `Sms` whether define in suite feature settings. expression like
// `Sms = SmsJuhe` is mean whether `Sms` define in suite feature settings and value
// is `SmsJuhe``
func CfgIf(expression string) bool {
	return features.CfgIf(expression)
}

func GetOssDomain() string {
	uri := "https://"
	if CfgIf("AliOSS") {
		return uri + AliOSSSetting.Domain + "/"
	} else if CfgIf("COS") {
		return uri + COSSetting.Domain + "/"
	} else if CfgIf("HuaweiOBS") {
		return uri + HuaweiOBSSetting.Domain + "/"
	} else if CfgIf("MinIO") {
		if !MinIOSetting.Secure {
			uri = "http://"
		}
		return uri + MinIOSetting.Domain + "/" + MinIOSetting.Bucket + "/"
	} else if CfgIf("S3") {
		if !S3Setting.Secure {
			uri = "http://"
		}
		// TODO: will not work well need test in real world
		return uri + S3Setting.Domain + "/" + S3Setting.Bucket + "/"
	} else if CfgIf("LocalOSS") {
		if !LocalOSSSetting.Secure {
			uri = "http://"
		}
		return uri + LocalOSSSetting.Domain + "/oss/" + LocalOSSSetting.Bucket + "/"
	}
	return uri + AliOSSSetting.Domain + "/"
}
