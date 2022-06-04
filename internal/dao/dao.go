package dao

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"gorm.io/gorm"
)

var (
	_ core.DataService = (*dataServant)(nil)
)

type dataServant struct {
	engine *gorm.DB
	zinc   *zinc.ZincClient
}

func NewDataService(engine *gorm.DB, zinc *zinc.ZincClient) core.DataService {
	return &dataServant{
		engine: engine,
		zinc:   zinc,
	}
}
