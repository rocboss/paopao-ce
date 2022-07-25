package storage

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func MustAliossService() (core.ObjectStorageService, core.VersionInfo) {
	client, err := oss.New(conf.AliOSSSetting.Endpoint, conf.AliOSSSetting.AccessKeyID, conf.AliOSSSetting.AccessKeySecret)
	if err != nil {
		logrus.Fatalf("storage.MustAliossService create client err: %s", err)
	}

	bucket, err := client.Bucket(conf.AliOSSSetting.Bucket)
	if err != nil {
		logrus.Fatalf("storage.MustAliossService create bucket err: %v", err)
	}

	obj := &aliossServant{
		bucket:             bucket,
		domain:             conf.GetOssDomain(),
		retainInDays:       time.Duration(conf.ObjectStorage.RetainInDays) * time.Hour * 24,
		retainUntilDate:    time.Date(2049, time.December, 1, 12, 0, 0, 0, time.UTC),
		allowPersistObject: conf.CfgIf("PersistObject"),
	}
	return obj, obj
}

func NewCosService() (core.ObjectStorageService, core.VersionInfo) {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", conf.COSSetting.Bucket, conf.COSSetting.Region))
	su, _ := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", conf.COSSetting.Region))

	client := cos.NewClient(&cos.BaseURL{BucketURL: u, ServiceURL: su}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.COSSetting.SecretID,
			SecretKey: conf.COSSetting.SecretKey,
		},
	})

	obj := &cosServant{
		client: client,
		domain: conf.GetOssDomain(),
	}
	return obj, obj
}

func MustHuaweiobsService() (core.ObjectStorageService, core.VersionInfo) {
	s := conf.HuaweiOBSSetting
	client, err := obs.New(s.AccessKey, s.SecretKey, s.Endpoint)
	if err != nil {
		logrus.Fatalf("storage.MustHuaweiobsService create huawei obs client failed: %s", err)
	}

	retainUntilDays := time.Until(time.Date(2049, time.December, 1, 12, 0, 0, 0, time.UTC)) / (24 * time.Hour)
	obj := &huaweiobsServant{
		client:             client,
		bucket:             s.Bucket,
		domain:             conf.GetOssDomain(),
		retainInDays:       int64(conf.ObjectStorage.RetainInDays),
		retainUntilDays:    strconv.FormatInt(int64(retainUntilDays), 10),
		allowPersistObject: conf.CfgIf("PersistObject"),
	}
	return obj, obj
}

func MustLocalossService() (core.ObjectStorageService, core.VersionInfo) {
	savePath, err := filepath.Abs(conf.LocalOSSSetting.SavePath)
	if err != nil {
		logrus.Fatalf("storage.MustLocalossService get localOSS save path err: %v", err)
	}

	obj := &localossServant{
		savePath: savePath + "/" + conf.LocalOSSSetting.Bucket + "/",
		domain:   conf.GetOssDomain(),
	}
	return obj, obj
}

func MustMinioService() (core.ObjectStorageService, core.VersionInfo) {
	// Initialize minio client object.
	client, err := minio.New(conf.MinIOSetting.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.MinIOSetting.AccessKey, conf.MinIOSetting.SecretKey, ""),
		Secure: conf.MinIOSetting.Secure,
	})
	if err != nil {
		logrus.Fatalf("storage.MustMinioService create client failed: %s", err)
	}

	ms := &minioServant{
		client:             client,
		bucket:             conf.MinIOSetting.Bucket,
		domain:             conf.GetOssDomain(),
		retainInDays:       time.Duration(conf.ObjectStorage.RetainInDays) * time.Hour * 24,
		retainUntilDate:    time.Date(2049, time.December, 1, 12, 0, 0, 0, time.UTC),
		allowPersistObject: conf.CfgIf("PersistObject"),
	}

	return ms, ms
}

func MustS3Service() (core.ObjectStorageService, core.VersionInfo) {
	// Initialize s3 client object use minio-go.
	client, err := minio.New(conf.S3Setting.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.S3Setting.AccessKey, conf.S3Setting.SecretKey, ""),
		Secure: conf.S3Setting.Secure,
	})
	if err != nil {
		logrus.Fatalf("storage.MustS3Service create client failed: %s", err)
	}

	s3 := &s3Servant{
		client:             client,
		bucket:             conf.MinIOSetting.Bucket,
		domain:             conf.GetOssDomain(),
		retainInDays:       time.Duration(conf.ObjectStorage.RetainInDays) * time.Hour * 24,
		retainUntilDate:    time.Date(2049, time.December, 1, 12, 0, 0, 0, time.UTC),
		allowPersistObject: conf.CfgIf("PersistObject"),
	}

	return s3, s3
}
