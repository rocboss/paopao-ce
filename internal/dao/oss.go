package dao

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/minio/minio-go/v7"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

var (
	_ core.ObjectStorageService = (*aliossServant)(nil)
	_ core.ObjectStorageService = (*minioServant)(nil)
	_ core.ObjectStorageService = (*s3Servant)(nil)
	_ core.ObjectStorageService = (*localossServant)(nil)
)

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

func NewObjectStorageService() (oss core.ObjectStorageService) {
	var v versionInfo
	if conf.CfgIf("AliOSS") {
		oss, v = newAliossServent()
	} else if conf.CfgIf("MinIO") {
		oss, v = newMinioServeant()
	} else if conf.CfgIf("S3") {
		oss, v = newS3Servent()
		logrus.Infof("use S3 as object storage by version %s", v.version())
		return
	} else if conf.CfgIf("LocalOSS") {
		oss, v = newLocalossServent()
	} else {
		// default use AliOSS as object storage service
		oss, v = newAliossServent()
		logrus.Infof("use default AliOSS as object storage by version %s", v.version())
		return
	}
	logrus.Infof("use %s as object storage by version %s", v.name(), v.version())
	return
}
