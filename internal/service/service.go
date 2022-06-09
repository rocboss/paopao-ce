package service

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/pkg/zinc"
)

var (
	ds                 core.DataService
	attachmentChecker  core.AttachmentCheckService
	DisablePhoneVerify bool
)

func Initialize() {
	zincClient := zinc.NewClient(conf.ZincSetting)
	ds = dao.NewDataService(conf.DBEngine, zincClient)

	attachmentChecker = dao.NewAttachmentCheckerService()
	DisablePhoneVerify = !conf.CfgIf("Sms")
}
