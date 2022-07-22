package service

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
)

var (
	ds                 core.DataService
	ts                 core.TweetSearchService
	oss                core.ObjectStorageService
	DisablePhoneVerify bool
)

func Initialize() {
	ds = dao.DataService()
	ts = dao.TweetSearchService()
	oss = dao.ObjectStorageService()
	DisablePhoneVerify = !conf.CfgIf("Sms")
}
