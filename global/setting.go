package global

import (
	"sync"

	"github.com/rocboss/paopao-ce/pkg/setting"
	"github.com/sirupsen/logrus"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RedisSetting    *setting.RedisSettingS
	SearchSetting   *setting.SearchSettingS
	AliossSetting   *setting.AliossSettingS
	JWTSetting      *setting.JWTSettingS
	LoggerSetting   *setting.LoggerSettingS
	Logger          *logrus.Logger
	Mutex           *sync.Mutex
)
