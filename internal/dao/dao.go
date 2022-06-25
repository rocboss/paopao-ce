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
	useCacheIndex     bool
	cacheIndex        core.CacheIndexService
	ams               core.AuthorizationManageService
	engine            *gorm.DB
	getIndexPostsFunc indexPostsFunc
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
	ds.useCacheIndex = true
	if conf.CfgIf("SimpleCacheIndex") {
		ds.getIndexPostsFunc = ds.simpleCacheIndexGetPosts
		ds.cacheIndex = newSimpleCacheIndexServant(ds.simpleCacheIndexGetPosts)
	} else if conf.CfgIf("BigCacheIndex") {
		ds.getIndexPostsFunc = ds.getIndexPosts
		ds.cacheIndex = newBigCacheIndexServant(ds.getIndexPosts)
	} else {
		ds.useCacheIndex = false
	}

	if ds.useCacheIndex {
		logrus.Infof("use %s as cache index service by version: %s", ds.cacheIndex.Name(), ds.cacheIndex.Version())
	}

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
