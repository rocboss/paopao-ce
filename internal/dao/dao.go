package dao

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	_ core.DataService                = (*dataServant)(nil)
	_ core.AttachmentCheckService     = (*attachmentCheckServant)(nil)
	_ core.AuthorizationManageService = (*simpleAuthorizationManageService)(nil)
)

type dataServant struct {
	core.CacheIndexService

	ams    core.AuthorizationManageService
	engine *gorm.DB
}

type simpleAuthorizationManageService struct {
	db *gorm.DB
}

type attachmentCheckServant struct {
	domain string
}

func NewDataService() core.DataService {
	ds := &dataServant{
		engine: conf.DBEngine,
		ams:    NewAuthorizationManageService(),
	}

	// initialize CacheIndex if needed
	if conf.CfgIf("SimpleCacheIndex") {
		ds.CacheIndexService = newSimpleCacheIndexServant(ds.simpleCacheIndexGetPosts)
	} else if conf.CfgIf("BigCacheIndex") {
		ds.CacheIndexService = newBigCacheIndexServant(ds.getIndexPosts)
	} else {
		ds.CacheIndexService = newNoneCacheIndexServant(ds.getIndexPosts)
	}
	logrus.Infof("use %s as cache index service by version: %s", ds.CacheIndexService.Name(), ds.CacheIndexService.Version())

	return ds
}

func NewAuthorizationManageService() core.AuthorizationManageService {
	return newSimpleAuthorizationManageService()
}

func NewAttachmentCheckService() core.AttachmentCheckService {
	return &attachmentCheckServant{
		domain: getOssDomain(),
	}
}
