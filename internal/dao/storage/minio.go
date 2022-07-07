package storage

import (
	"context"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/minio/minio-go/v7"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

var (
	_ core.ObjectStorageService = (*minioServant)(nil)
	_ core.VersionInfo          = (*minioServant)(nil)
)

type minioServant struct {
	client *minio.Client
	bucket string
	domain string
}

type s3Servant = minioServant

func (s *minioServant) Name() string {
	return "MinIO"
}

func (s *minioServant) Version() *semver.Version {
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
