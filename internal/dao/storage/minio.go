// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
	_ core.OssCreateService     = (*minioCreateServant)(nil)
	_ core.OssCreateService     = (*minioCreateRetentionServant)(nil)
	_ core.OssCreateService     = (*minioCreateTempDirServant)(nil)
	_ core.VersionInfo          = (*minioServant)(nil)
)

type minioCreateServant struct {
	client *minio.Client
	bucket string
	domain string
}

type minioCreateRetentionServant struct {
	client          *minio.Client
	bucket          string
	domain          string
	retainInDays    time.Duration
	retainUntilDate time.Time
}

type minioCreateTempDirServant struct {
	client  *minio.Client
	bucket  string
	domain  string
	tempDir string
}

type minioServant struct {
	core.OssCreateService

	client *minio.Client
	bucket string
	domain string
}

type s3Servant = minioServant

func (s *minioCreateServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, _persistance bool) (string, error) {
	opts := minio.PutObjectOptions{ContentType: contentType}
	_, err := s.client.PutObject(context.Background(), s.bucket, objectKey, reader, objectSize, opts)
	if err != nil {
		return "", err
	}
	return s.domain + objectKey, nil
}

func (s *minioCreateServant) PersistObject(_objectKey string) error {
	// empty
	return nil
}

func (s *minioCreateRetentionServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, persistance bool) (string, error) {
	opts := minio.PutObjectOptions{ContentType: contentType}
	if !persistance {
		opts.Mode = minio.Governance
		opts.RetainUntilDate = time.Now().Add(s.retainInDays)
	}
	_, err := s.client.PutObject(context.Background(), s.bucket, objectKey, reader, objectSize, opts)
	if err != nil {
		return "", err
	}
	return s.domain + objectKey, nil
}

func (s *minioCreateRetentionServant) PersistObject(objectKey string) error {
	retentionMode := minio.Governance
	return s.client.PutObjectRetention(context.Background(), s.bucket, objectKey, minio.PutObjectRetentionOptions{
		Mode:            &retentionMode,
		RetainUntilDate: &s.retainUntilDate,
	})
}

func (s *minioCreateTempDirServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, persistance bool) (string, error) {
	opts := minio.PutObjectOptions{ContentType: contentType}
	objectName := objectKey
	if !persistance {
		objectName = s.tempDir + objectKey
	}
	_, err := s.client.PutObject(context.Background(), s.bucket, objectName, reader, objectSize, opts)
	if err != nil {
		return "", err
	}
	return s.domain + objectKey, nil
}

func (s *minioCreateTempDirServant) PersistObject(objectKey string) error {
	_, err := s.client.StatObject(context.Background(), s.bucket, objectKey, minio.StatObjectOptions{})
	if err == nil {
		logrus.Debugf("object exist so do nothing objectKey: %s", objectKey)
		return nil
	}

	tmpObjKey := s.tempDir + objectKey
	src := minio.CopySrcOptions{
		Bucket: s.bucket,
		Object: tmpObjKey,
	}
	dst := minio.CopyDestOptions{
		Bucket: s.bucket,
		Object: objectKey,
	}
	if _, err = s.client.CopyObject(context.Background(), dst, src); err != nil {
		return err
	}

	return s.client.RemoveObject(context.Background(), s.bucket, tmpObjKey, minio.RemoveObjectOptions{ForceDelete: true})
}

func (s *minioServant) DeleteObject(objectKey string) error {
	return s.client.RemoveObject(context.Background(), s.bucket, objectKey, minio.RemoveObjectOptions{ForceDelete: true})
}

func (s *minioServant) DeleteObjects(objectKeys []string) (err error) {
	objectsCh := make(chan minio.ObjectInfo, len(objectKeys))

	resCh := s.client.RemoveObjects(context.Background(), s.bucket, objectsCh, minio.RemoveObjectsOptions{})
	for _, objectKey := range objectKeys {
		objectsCh <- minio.ObjectInfo{
			Key: objectKey,
		}
	}

	// 宽松处理所有错误，只记录最后一次发生的错误
	for result := range resCh {
		if result.Err != nil {
			err = result.Err
		}
	}
	return
}

func (s *minioServant) IsObjectExist(objectKey string) (bool, error) {
	_, err := s.client.StatObject(context.Background(), s.bucket, objectKey, minio.StatObjectOptions{})
	if err != nil {
		return false, err
	}
	return true, nil
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

func (s *minioServant) Name() string {
	return "MinIO"
}

func (s *minioServant) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}
