package dao

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/minio/minio-go/v7"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/zinc"
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
	return &dataServant{
		engine: engine,
		zinc:   zinc,
	}
}

func NewObjectStorageService() (oss core.ObjectStorageService) {
	if conf.CfgIf("AliOSS") {
		oss = newAliossServent()
		conf.Logger.Infoln("use AliOSS as object storage")
	} else if conf.CfgIf("MinIO") {
		oss = newMinioServeant()
		conf.Logger.Infoln("use MinIO as object storage")
	} else if conf.CfgIf("S3") {
		oss = newS3Servent()
		conf.Logger.Infoln("use S3 as object storage")
	} else if conf.CfgIf("LocalOSS") {
		oss = newLocalossServent()
		conf.Logger.Infoln("use LocalOSS as object storage")
	} else {
		// default use AliOSS
		oss = newAliossServent()
		conf.Logger.Infoln("use default AliOSS as object storage")
	}
	return
}

func NewAttachmentCheckerService() core.AttachmentCheckService {
	return &attachmentCheckServant{
		domain: getOssDomain(),
	}
}
