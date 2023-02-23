// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var (
	_ core.ObjectStorageService = (*cosServant)(nil)
	_ core.OssCreateService     = (*cosCreateServant)(nil)
	_ core.OssCreateService     = (*cosCreateTempDirServant)(nil)
	_ core.VersionInfo          = (*cosServant)(nil)
)

type cosCreateServant struct {
	client *cos.Client
	domain string
}

type cosCreateTempDirServant struct {
	client    *cos.Client
	domain    string
	bucketUrl string
	tempDir   string
}

type cosServant struct {
	core.OssCreateService

	client *cos.Client
	domain string
}

func (s *cosCreateServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, _persistance bool) (string, error) {
	_, err := s.client.Object.Put(context.Background(), objectKey, reader, &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType:   contentType,
			ContentLength: objectSize,
		},
	})
	if err != nil {
		return "", err
	}
	return s.domain + objectKey, nil
}

func (s *cosCreateServant) PersistObject(_objectKey string) error {
	// empty
	return nil
}

func (s *cosCreateTempDirServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, persistance bool) (string, error) {
	objectName := objectKey
	if !persistance {
		objectName = s.tempDir + objectKey
	}
	_, err := s.client.Object.Put(context.Background(), objectName, reader, &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType:   contentType,
			ContentLength: objectSize,
		},
	})
	if err != nil {
		return "", err
	}
	return s.domain + objectKey, nil
}

func (s *cosCreateTempDirServant) PersistObject(objectKey string) error {
	exsit, err := s.client.Object.IsExist(context.Background(), objectKey)
	if err != nil {
		return err
	}
	if exsit {
		logrus.Debugf("object exist so do nothing objectKey: %s", objectKey)
		return nil
	}
	_, _, err = s.client.Object.Copy(context.Background(), objectKey, s.bucketUrl+s.tempDir+objectKey, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Object.Delete(context.Background(), s.tempDir+objectKey, nil)
	return err
}

func (s *cosServant) DeleteObject(objectKey string) error {
	_, err := s.client.Object.Delete(context.Background(), objectKey)
	return err
}

func (s *cosServant) DeleteObjects(objectKeys []string) error {
	obs := []cos.Object{}
	for _, v := range objectKeys {
		obs = append(obs, cos.Object{Key: v})
	}
	_, _, err := s.client.Object.DeleteMulti(context.Background(), &cos.ObjectDeleteMultiOptions{
		Objects: obs,
	})
	return err
}

func (s *cosServant) IsObjectExist(objectKey string) (bool, error) {
	return s.client.Object.IsExist(context.Background(), objectKey)
}

func (s *cosServant) SignURL(objectKey string, expiredInSec int64) (string, error) {
	signedURL, err := s.client.Object.GetPresignedURL(context.Background(),
		http.MethodGet, objectKey, conf.COSSetting.SecretID, conf.COSSetting.SecretKey, time.Second*time.Duration(expiredInSec), nil)

	if err != nil {
		logrus.Errorf("client.SignURL err: %v", err)
		return "", err
	}

	epath, err := url.PathUnescape(signedURL.Path)
	if err != nil {
		logrus.Errorf("url.PathUnescape err: %v", err)
		return "", err
	}

	signedURL.Path, signedURL.RawPath = epath, epath
	return signedURL.String(), nil
}

func (s *cosServant) ObjectURL(objetKey string) string {
	return s.domain + objetKey
}

func (s *cosServant) ObjectKey(objectUrl string) string {
	return strings.Replace(objectUrl, s.domain, "", -1)
}

func (s *cosServant) Name() string {
	return "COS"
}

func (s *cosServant) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}
