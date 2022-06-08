package dao

import (
	"context"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rocboss/paopao-ce/global"
)

func newMinioServeant() *minioServant {
	// Initialize minio client object.
	client, err := minio.New(global.MinIOSetting.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.MinIOSetting.AccessKey, global.MinIOSetting.SecretKey, ""),
		Secure: global.MinIOSetting.Secure,
	})
	if err != nil {
		global.Logger.Fatalf("minio.New err: %v", err)
	}
	return &minioServant{
		client: client,
		bucket: global.MinIOSetting.Bucket,
		domain: getOssDomain(),
	}
}

func (s *minioServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	uploadInfo, err := s.client.PutObject(context.Background(), s.bucket, objectKey, reader, objectSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}
	global.Logger.Infoln("Successfully uploaded bytes: ", uploadInfo)
	return s.domain + objectKey, nil
}
func (s *minioServant) SignURL(objectKey string, expiredInSec int64) (string, error) {
	// TODO: Set request parameters for content-disposition.
	reqParams := make(url.Values)
	signedURL, err := s.client.PresignedGetObject(context.Background(), s.bucket, objectKey, time.Duration(expiredInSec)*time.Second, reqParams)
	if err != nil {
		return "", err
	}
	return signedURL.String(), nil
}

func (s *minioServant) ObjectURL(objetKey string) string {
	return s.domain + objetKey
}

func (s *minioServant) ObjectKey(objectUrl string) string {
	return strings.Replace(objectUrl, s.domain, "", -1)
}
