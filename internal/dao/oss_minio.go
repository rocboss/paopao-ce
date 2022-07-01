package dao

import (
	"context"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

func newMinioServeant() (*minioServant, versionInfo) {
	// Initialize minio client object.
	client, err := minio.New(conf.MinIOSetting.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.MinIOSetting.AccessKey, conf.MinIOSetting.SecretKey, ""),
		Secure: conf.MinIOSetting.Secure,
	})
	if err != nil {
		logrus.Fatalf("minio.New err: %v", err)
	}

	obj := &minioServant{
		client: client,
		bucket: conf.MinIOSetting.Bucket,
		domain: getOssDomain(),
	}
	return obj, obj
}

func (s *minioServant) name() string {
	return "MinIO"
}

func (s *minioServant) version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *minioServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	uploadInfo, err := s.client.PutObject(context.Background(), s.bucket, objectKey, reader, objectSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}
	logrus.Infoln("Successfully uploaded bytes: ", uploadInfo)
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
