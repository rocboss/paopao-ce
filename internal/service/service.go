package service

import (
	"github.com/rocboss/paopao-ce/internal/dao"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"gorm.io/gorm"
)

var (
	myDao *dao.Dao
)

func Initialize(engine *gorm.DB, client *zinc.ZincClient) {
	myDao = dao.New(engine, client)
}
