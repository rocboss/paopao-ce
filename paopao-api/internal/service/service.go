package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-api/global"
	"github.com/rocboss/paopao-api/internal/dao"
	"github.com/rocboss/paopao-api/pkg/zinc"
)

type Service struct {
	ctx *gin.Context
	dao *dao.Dao
}

func New(ctx *gin.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine, &zinc.ZincClient{
		ZincClientConfig: &zinc.ZincClientConfig{
			ZincHost:     global.SearchSetting.ZincHost,
			ZincUser:     global.SearchSetting.ZincUser,
			ZincPassword: global.SearchSetting.ZincPassword,
		},
	})

	return svc
}
