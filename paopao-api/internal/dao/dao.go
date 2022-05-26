package dao

import (
	"github.com/rocboss/paopao-api/pkg/zinc"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
	zinc   *zinc.ZincClient
}

func New(engine *gorm.DB, zinc *zinc.ZincClient) *Dao {
	return &Dao{
		engine: engine,
		zinc:   zinc,
	}
}
