package dao

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

func newS3Servent() *s3Servant {
	// Initialize s3 client object use minio-go.
	client, err := minio.New(conf.S3Setting.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.S3Setting.AccessKey, conf.S3Setting.SecretKey, ""),
		Secure: conf.S3Setting.Secure,
	})
	if err != nil {
		logrus.Fatalf("s3.New err: %v", err)
	}
	return &s3Servant{
		client: client,
		bucket: conf.MinIOSetting.Bucket,
		domain: getOssDomain(),
	}
}
