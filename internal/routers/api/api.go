package api

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
)

var (
	objectStorage core.ObjectStorageService
)

func Initialize() {
	objectStorage = dao.ObjectStorageService()
}
