package storage

import (
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

var (
	_ core.ObjectStorageService = (*aliossServant)(nil)
	_ core.VersionInfo          = (*aliossServant)(nil)
)

type aliossServant struct {
	bucket             *oss.Bucket
	domain             string
	retainInDays       time.Duration
	retainUntilDate    time.Time
	allowPersistObject bool
}

func (s *aliossServant) Name() string {
	return "AliOSS"
}

func (s *aliossServant) Version() *semver.Version {
	return semver.MustParse("v0.2.0")
}

func (s *aliossServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, persistance bool) (string, error) {
	options := []oss.Option{
		oss.ContentLength(objectSize),
		oss.ContentType(contentType),
	}
	if s.allowPersistObject && !persistance {
		options = append(options, oss.Expires(time.Now().Add(s.retainInDays)))
	}
	err := s.bucket.PutObject(objectKey, reader, options...)
	if err != nil {
		return "", err
	}
	return s.domain + objectKey, nil
}

func (s *aliossServant) PersistObject(objectKey string) error {
	if !s.allowPersistObject {
		return nil
	}
	return s.bucket.SetObjectMeta(objectKey, oss.Expires(s.retainUntilDate))
}

func (s *aliossServant) DeleteObject(objectKey string) error {
	return s.bucket.DeleteObject(objectKey)
}

func (s *aliossServant) DeleteObjects(objectKeys []string) error {
	_, err := s.bucket.DeleteObjects(objectKeys)
	return err
}

func (s *aliossServant) IsObjectExist(objectKey string) (bool, error) {
	return s.bucket.IsObjectExist(objectKey)
}

func (s *aliossServant) SignURL(objectKey string, expiredInSec int64) (string, error) {
	signedURL, err := s.bucket.SignURL(objectKey, oss.HTTPGet, expiredInSec)
	if err != nil {
		logrus.Errorf("client.SignURL err: %v", err)
		return "", err
	}

	ur, err := url.Parse(signedURL)
	if err != nil {
		logrus.Errorf("url.Parse err: %v", err)
		return "", err
	}

	epath, err := url.PathUnescape(ur.Path)
	if err != nil {
		logrus.Errorf("url.PathUnescape err: %v", err)
		return "", err
	}

	ur.Path, ur.RawPath = epath, epath
	return ur.String(), nil
}

func (s *aliossServant) ObjectURL(objetKey string) string {
	return s.domain + objetKey
}

func (s *aliossServant) ObjectKey(objectUrl string) string {
	return strings.Replace(objectUrl, s.domain, "", -1)
}
