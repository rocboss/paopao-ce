package storage

import (
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

func NewAliossService() (core.ObjectStorageService, core.VersionInfo) {
	client, err := oss.New(conf.AliOSSSetting.Endpoint, conf.AliOSSSetting.AccessKeyID, conf.AliOSSSetting.AccessKeySecret)
	if err != nil {
		logrus.Fatalf("alioss.New err: %v", err)
	}

	bucket, err := client.Bucket(conf.AliOSSSetting.Bucket)
	if err != nil {
		logrus.Fatalf("client.Bucket err: %v", err)
	}

	obj := &aliossServant{
		bucket: bucket,
		domain: conf.GetOssDomain(),
	}
	return obj, obj
}

func NewLocalossService() (core.ObjectStorageService, core.VersionInfo) {
	savePath, err := filepath.Abs(conf.LocalOSSSetting.SavePath)
	if err != nil {
		logrus.Fatalf("get localOSS save path err: %v", err)
	}

	obj := &localossServant{
		savePath: savePath + "/" + conf.LocalOSSSetting.Bucket + "/",
		domain:   conf.GetOssDomain(),
	}
	return obj, obj
}

func NewMinioService() (core.ObjectStorageService, core.VersionInfo) {
	// Initialize minio client object.
	client, err := minio.New(conf.MinIOSetting.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.MinIOSetting.AccessKey, conf.MinIOSetting.SecretKey, ""),
		Secure: conf.MinIOSetting.Secure,
	})
	if err != nil {
		logrus.Fatalf("minio.New err: %v", err)
	}

	ms := &minioServant{
		client: client,
		bucket: conf.MinIOSetting.Bucket,
		domain: conf.GetOssDomain(),
	}
	return ms, ms
}

func NewS3Service() (core.ObjectStorageService, core.VersionInfo) {
	// Initialize s3 client object use minio-go.
	client, err := minio.New(conf.S3Setting.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.S3Setting.AccessKey, conf.S3Setting.SecretKey, ""),
		Secure: conf.S3Setting.Secure,
	})
	if err != nil {
		logrus.Fatalf("s3.New err: %v", err)
	}

	obj := &s3Servant{
		client: client,
		bucket: conf.MinIOSetting.Bucket,
		domain: conf.GetOssDomain(),
	}
	return obj, obj
}
