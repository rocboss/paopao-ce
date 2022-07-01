package service

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
)

var (
	ds                 core.DataService
	ts                 core.TweetSearchService
	DisablePhoneVerify bool
)

func Initialize() {
	ds = dao.NewDataService()
	ts = dao.NewTweetSearchService()
	DisablePhoneVerify = !conf.CfgIf("Sms")
}
