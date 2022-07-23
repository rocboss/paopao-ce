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
	_ core.VersionInfo          = (*cosServant)(nil)
)

type cosServant struct {
	client *cos.Client
	domain string
}

func (s *cosServant) Name() string {
	return "COS"
}

func (s *cosServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *cosServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	_, err := s.client.Object.Put(context.Background(), objectKey, reader, &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: contentType,
		},
	})
	if err != nil {
		return "", err
	}
	return s.domain + objectKey, nil
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
