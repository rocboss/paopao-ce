package dao

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rocboss/paopao-ce/global"
)

func newS3Servent() *s3Servant {
	// Initialize s3 client object use minio-go.
	client, err := minio.New(global.S3Setting.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.S3Setting.AccessKey, global.S3Setting.SecretKey, ""),
		Secure: global.S3Setting.Secure,
	})
	if err != nil {
		global.Logger.Fatalf("s3.New err: %v", err)
	}
	return &s3Servant{
		client: client,
		bucket: global.MinIOSetting.Bucket,
		domain: getOssDomain(),
	}
}
