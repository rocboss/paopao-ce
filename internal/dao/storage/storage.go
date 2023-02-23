// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"

	"github.com/alimy/cfg"
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
		logrus.Fatalf("storage.MustAliossService create bucket err: %s", err)
	}

	domain := conf.GetOssDomain()
	var cs core.OssCreateService
	if cfg.If("OSS:TempDir") {
		cs = &aliossCreateTempDirServant{
			bucket:  bucket,
			domain:  domain,
			tempDir: conf.ObjectStorage.TempDirSlash(),
		}
		logrus.Debugln("use OSS:TempDir feature")
	} else if cfg.If("OSS:Retention") {
		cs = &aliossCreateRetentionServant{
			bucket:          bucket,
			domain:          domain,
			retainInDays:    time.Duration(conf.ObjectStorage.RetainInDays) * time.Hour * 24,
			retainUntilDate: time.Date(2049, time.December, 1, 12, 0, 0, 0, time.UTC),
		}
		logrus.Debugln("use OSS:Retention feature")
	} else {
		cs = &aliossCreateServant{
			bucket: bucket,
			domain: domain,
		}
		logrus.Debugln("use OSS:Direct feature")
	}

	obj := &aliossServant{
		OssCreateService: cs,
		bucket:           bucket,
		domain:           conf.GetOssDomain(),
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

	domain := conf.GetOssDomain()
	var cs core.OssCreateService
	if cfg.If("OSS:TempDir") {
		cs = &cosCreateTempDirServant{
			client:    client,
			domain:    domain,
			bucketUrl: fmt.Sprintf("%s.cos.%s.myqcloud.com/", conf.COSSetting.Bucket, conf.COSSetting.Region),
			tempDir:   conf.ObjectStorage.TempDirSlash(),
		}
		logrus.Debugln("use OSS:TempDir feature")
	} else {
		cs = &cosCreateServant{
			client: client,
			domain: domain,
		}
		logrus.Debugln("use OSS:Direct feature")
	}

	obj := &cosServant{
		OssCreateService: cs,
		client:           client,
		domain:           domain,
	}
	return obj, obj
}

func MustHuaweiobsService() (core.ObjectStorageService, core.VersionInfo) {
	s := conf.HuaweiOBSSetting
	client, err := obs.New(s.AccessKey, s.SecretKey, s.Endpoint)
	if err != nil {
		logrus.Fatalf("storage.MustHuaweiobsService create huawei obs client failed: %s", err)
	}

	domain := conf.GetOssDomain()
	var cs core.OssCreateService
	if cfg.If("OSS:TempDir") {
		cs = &hwobsCreateTempDirServant{
			client:  client,
			bucket:  s.Bucket,
			domain:  domain,
			tempDir: conf.ObjectStorage.TempDirSlash(),
		}
		logrus.Debugln("use OSS:TempDir feature")
	} else if cfg.If("OSS:Retention") {
		retainUntilDays := time.Until(time.Date(2049, time.December, 1, 12, 0, 0, 0, time.UTC)) / (24 * time.Hour)
		cs = &hwobsCreateRetentionServant{
			client:          client,
			bucket:          s.Bucket,
			domain:          domain,
			retainInDays:    int64(conf.ObjectStorage.RetainInDays),
			retainUntilDays: strconv.FormatInt(int64(retainUntilDays), 10),
		}
		logrus.Debugln("use OSS:Retention feature")
	} else {
		cs = &hwobsCreateServant{
			client: client,
			bucket: s.Bucket,
			domain: domain,
		}
		logrus.Debugln("use OSS:Direct feature")
	}

	obj := &huaweiobsServant{
		OssCreateService: cs,
		client:           client,
		bucket:           s.Bucket,
		domain:           domain,
	}
	return obj, obj
}

func MustLocalossService() (core.ObjectStorageService, core.VersionInfo) {
	savePath, err := filepath.Abs(conf.LocalOSSSetting.SavePath)
	if err != nil {
		logrus.Fatalf("storage.MustLocalossService get localOSS save path err: %s", err)
	}

	domain := conf.GetOssDomain()
	savePath = savePath + "/" + conf.LocalOSSSetting.Bucket + "/"
	var cs core.OssCreateService
	if cfg.If("OSS:TempDir") {
		cs = &localossCreateTempDirServant{
			savePath: savePath,
			domain:   domain,
			tempDir:  conf.ObjectStorage.TempDirSlash(),
		}
		logrus.Debugln("use OSS:TempDir feature")
	} else {
		cs = &localossCreateServant{
			savePath: savePath,
			domain:   domain,
		}
		logrus.Debugln("use OSS:Direct feature")
	}

	obj := &localossServant{
		OssCreateService: cs,
		savePath:         savePath,
		domain:           domain,
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

	domain := conf.GetOssDomain()
	var cs core.OssCreateService
	if cfg.If("OSS:TempDir") {
		cs = &minioCreateTempDirServant{
			client:  client,
			bucket:  conf.MinIOSetting.Bucket,
			domain:  domain,
			tempDir: conf.ObjectStorage.TempDirSlash(),
		}
		logrus.Debugln("use OSS:TempDir feature")
	} else if cfg.If("OSS:Retention") {
		cs = &minioCreateRetentionServant{
			client:          client,
			bucket:          conf.MinIOSetting.Bucket,
			domain:          domain,
			retainInDays:    time.Duration(conf.ObjectStorage.RetainInDays) * time.Hour * 24,
			retainUntilDate: time.Date(2049, time.December, 1, 12, 0, 0, 0, time.UTC),
		}
		logrus.Debugln("use OSS:Retention feature")
	} else {
		cs = &minioCreateServant{
			client: client,
			bucket: conf.MinIOSetting.Bucket,
			domain: domain,
		}
		logrus.Debugln("use OSS:Direct feature")
	}

	obj := &minioServant{
		OssCreateService: cs,
		client:           client,
		bucket:           conf.MinIOSetting.Bucket,
		domain:           domain,
	}
	return obj, obj
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

	domain := conf.GetOssDomain()
	var cs core.OssCreateService
	if cfg.If("OSS:TempDir") {
		cs = &minioCreateTempDirServant{
			client:  client,
			bucket:  conf.S3Setting.Bucket,
			domain:  domain,
			tempDir: conf.ObjectStorage.TempDirSlash(),
		}
		logrus.Debugln("use OSS:TempDir feature")
	} else if cfg.If("OSS:Retention") {
		cs = &minioCreateRetentionServant{
			client:          client,
			bucket:          conf.S3Setting.Bucket,
			domain:          domain,
			retainInDays:    time.Duration(conf.ObjectStorage.RetainInDays) * time.Hour * 24,
			retainUntilDate: time.Date(2049, time.December, 1, 12, 0, 0, 0, time.UTC),
		}
		logrus.Debugln("use OSS:Retention feature")
	} else {
		cs = &minioCreateServant{
			client: client,
			bucket: conf.S3Setting.Bucket,
			domain: domain,
		}
		logrus.Debugln("use OSS:Direct feature")
	}

	obj := &s3Servant{
		OssCreateService: cs,
		client:           client,
		bucket:           conf.S3Setting.Bucket,
		domain:           domain,
	}
	return obj, obj
}
