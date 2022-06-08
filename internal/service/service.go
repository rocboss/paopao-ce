package service

import (
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"gorm.io/gorm"
)

var (
	ds                 core.DataService
	attachmentChecker  core.AttachmentCheckService
	DisablePhoneVerify bool
)

func Initialize(engine *gorm.DB, client *zinc.ZincClient) {
	ds = dao.NewDataService(engine, client)
	attachmentChecker = dao.NewAttachmentCheckerService()
	DisablePhoneVerify = !global.CfgIf("Sms")
}
