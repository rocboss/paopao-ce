package dao

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	_ core.DataService            = (*dataServant)(nil)
	_ core.AttachmentCheckService = (*attachmentCheckServant)(nil)
)

type dataServant struct {
	useCacheIndex bool
	cacheIndex    core.CacheIndexService

	engine *gorm.DB
	zinc   *zinc.ZincClient
}

type attachmentCheckServant struct {
	domain string
}

func NewDataService() core.DataService {
	client := zinc.NewClient(conf.ZincSetting)
	ds := &dataServant{
		engine: conf.DBEngine,
		zinc:   client,
	}

	// initialize CacheIndex if needed
	ds.useCacheIndex = true
	if conf.CfgIf("SimpleCacheIndex") {
		ds.cacheIndex = newSimpleCacheIndexServant(ds.getIndexPosts)
	} else if conf.CfgIf("BigCacheIndex") {
		ds.cacheIndex = newBigCacheIndexServant(ds.getIndexPosts)
	} else {
		ds.useCacheIndex = false
	}

	if ds.useCacheIndex {
		logrus.Infof("use %s as cache index service by version: %s", ds.cacheIndex.Name(), ds.cacheIndex.Version())
	}

	return ds
}

func NewAttachmentCheckService() core.AttachmentCheckService {
	return &attachmentCheckServant{
		domain: getOssDomain(),
	}
}
