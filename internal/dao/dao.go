package dao

import (
	"sync/atomic"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/minio/minio-go/v7"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/zinc"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	_ core.DataService            = (*dataServant)(nil)
	_ core.ObjectStorageService   = (*aliossServant)(nil)
	_ core.ObjectStorageService   = (*minioServant)(nil)
	_ core.ObjectStorageService   = (*s3Servant)(nil)
	_ core.ObjectStorageService   = (*localossServant)(nil)
	_ core.AttachmentCheckService = (*attachmentCheckServant)(nil)
)

type dataServant struct {
	engine *gorm.DB
	zinc   *zinc.ZincClient

	useCacheIndex   bool
	indexActionCh   chan indexActionT
	indexPosts      []*model.PostFormated
	atomicIndex     atomic.Value
	maxIndexSize    int
	checkTick       *time.Ticker
	expireIndexTick *time.Ticker
}

type localossServant struct {
	savePath string
	domain   string
}

type aliossServant struct {
	bucket *oss.Bucket
	domain string
}

type minioServant struct {
	client *minio.Client
	bucket string
	domain string
}

type s3Servant = minioServant

type attachmentCheckServant struct {
	domain string
}

func NewDataService(engine *gorm.DB, zinc *zinc.ZincClient) core.DataService {
	if !conf.CfgIf("CacheIndex") {
		return &dataServant{
			engine:        engine,
			zinc:          zinc,
			useCacheIndex: false,
		}
	}

	s := conf.CacheIndexSetting
	ds := &dataServant{
		engine:          engine,
		zinc:            zinc,
		useCacheIndex:   true,
		maxIndexSize:    conf.CacheIndexSetting.MaxIndexSize,
		indexPosts:      make([]*model.PostFormated, 0),
		indexActionCh:   make(chan indexActionT, 100),                                      // optimize: size need configure by custom
		checkTick:       time.NewTicker(time.Duration(s.CheckTickDuration) * time.Second),  // check whether need update index every 1 minute
		expireIndexTick: time.NewTicker(time.Duration(s.ExpireTickDuration) * time.Second), // force expire index every 5 minute
	}

	// start index posts
	ds.atomicIndex.Store(ds.indexPosts)
	go ds.startIndexPosts()

	return ds
}

func NewObjectStorageService() (oss core.ObjectStorageService) {
	if conf.CfgIf("AliOSS") {
		oss = newAliossServent()
		logrus.Infoln("use AliOSS as object storage")
	} else if conf.CfgIf("MinIO") {
		oss = newMinioServeant()
		logrus.Infoln("use MinIO as object storage")
	} else if conf.CfgIf("S3") {
		oss = newS3Servent()
		logrus.Infoln("use S3 as object storage")
	} else if conf.CfgIf("LocalOSS") {
		oss = newLocalossServent()
		logrus.Infoln("use LocalOSS as object storage")
	} else {
		// default use AliOSS
		oss = newAliossServent()
		logrus.Infoln("use default AliOSS as object storage")
	}
	return
}

func NewAttachmentCheckerService() core.AttachmentCheckService {
	return &attachmentCheckServant{
		domain: getOssDomain(),
	}
}
