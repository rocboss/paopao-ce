package api

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
)

var (
	objectStorage     core.ObjectStorageService
	attachmentChecker core.AttachmentCheckService
)

func Initialize() {
	objectStorage = dao.NewObjectStorageService()
	attachmentChecker = dao.NewAttachmentCheckerService()
}
