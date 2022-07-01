package dao

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	_ core.DataService                = (*dataServant)(nil)
	_ core.AuthorizationManageService = (*simpleAuthorizationManageService)(nil)
)

type dataServant struct {
	core.CacheIndexService
	core.AttachmentCheckService

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
		ams:                    newAuthorizationManageService(),
		AttachmentCheckService: NewAttachmentCheckService(),

		engine: conf.DBEngine,
	}

	// initialize CacheIndex if needed
	var v versionInfo
	if conf.CfgIf("SimpleCacheIndex") {
		ds.CacheIndexService, v = newSimpleCacheIndexServant(ds.simpleCacheIndexGetPosts)
	} else if conf.CfgIf("BigCacheIndex") {
		ds.CacheIndexService, v = newBigCacheIndexServant(ds.getIndexPosts)
	} else {
		ds.CacheIndexService, v = newNoneCacheIndexServant(ds.getIndexPosts)
	}
	logrus.Infof("use %s as cache index service by version: %s", v.name(), v.version())

	return ds
}

func NewAttachmentCheckService() core.AttachmentCheckService {
	return &attachmentCheckServant{
		domain: getOssDomain(),
	}
}

func newAuthorizationManageService() core.AuthorizationManageService {
	return newSimpleAuthorizationManageService()
}
