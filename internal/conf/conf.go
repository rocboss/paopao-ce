package conf

import (
	"log"
	"sync"
	"time"
)

var (
	loggerFileSetting *LoggerFileSettingS
	loggerZincSetting *LoggerZincSettingS
	databaseSetting   *DatabaseSetingS
	mysqlSetting      *MySQLSettingS
	postgresSetting   *PostgresSettingS
	sqlite3Setting    *Sqlite3SettingS
	redisSetting      *RedisSettingS
	features          *FeaturesSettingS

	ServerSetting           *ServerSettingS
	AppSetting              *AppSettingS
	SimpleCacheIndexSetting *SimpleCacheIndexSettingS
	BigCacheIndexSetting    *BigCacheIndexSettingS
	SmsJuheSetting          *SmsJuheSettings
	AlipaySetting           *AlipaySettingS
	ZincSetting             *ZincSettingS
	AliOSSSetting           *AliOSSSettingS
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
		"SimpleCacheIndex": &SimpleCacheIndexSetting,
		"BigCacheIndex":    &BigCacheIndexSetting,
		"Alipay":           &AlipaySetting,
		"SmsJuhe":          &SmsJuheSetting,
		"LoggerFile":       &loggerFileSetting,
		"LoggerZinc":       &loggerZincSetting,
		"Database":         &databaseSetting,
		"MySQL":            &mysqlSetting,
		"Postgres":         &postgresSetting,
		"Sqlite3":          &sqlite3Setting,
		"Zinc":             &ZincSetting,
		"Redis":            &redisSetting,
		"JWT":              &JWTSetting,
		"AliOSS":           &AliOSSSetting,
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

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
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
