package global

import (
	"sync"

	"github.com/rocboss/paopao-ce/pkg/setting"
	"github.com/sirupsen/logrus"
)

var (
	Features          *setting.FeaturesSettingS
	ServerSetting     *setting.ServerSettingS
	AppSetting        *setting.AppSettingS
	MySQLSetting      *setting.MySQLSettingS
	RedisSetting      *setting.RedisSettingS
	SmsJuheSetting    *setting.SmsJuheSettings
	AlipaySetting     *setting.AlipaySettingS
	ZincSetting       *setting.ZincSettingS
	AliOSSSetting     *setting.AliOSSSettingS
	MinIOSetting      *setting.MinIOSettingS
	S3Setting         *setting.S3SettingS
	LocalOSSSetting   *setting.LocalOSSSettingS
	JWTSetting        *setting.JWTSettingS
	LoggerFileSetting *setting.LoggerFileSettingS
	LoggerZincSetting *setting.LoggerZincSettingS
	Logger            *logrus.Logger
	Mutex             *sync.Mutex
)

func Cfg(key string) (string, bool) {
	return Features.Cfg(key)
}

func CfgIf(expression string) bool {
	return Features.CfgIf(expression)
}
