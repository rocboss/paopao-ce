package service

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"gorm.io/gorm"
)

var (
	ds core.DataService
)

func Initialize(engine *gorm.DB, client *zinc.ZincClient) {
	ds = dao.NewDataService(engine, client)
}
