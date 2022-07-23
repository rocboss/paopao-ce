package storage

import (
	"io"
	"net/url"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

var (
	_ core.ObjectStorageService = (*huaweiobsServant)(nil)
	_ core.VersionInfo          = (*huaweiobsServant)(nil)
)

type huaweiobsServant struct {
	client *obs.ObsClient
	bucket string
	domain string
}

func (s *huaweiobsServant) Name() string {
	return "HuaweiOBS"
}

func (s *huaweiobsServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *huaweiobsServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	input := &obs.PutObjectInput{}
	input.Bucket, input.Key, input.Body = s.bucket, objectKey, reader
	input.ContentType, input.ContentLength = contentType, objectSize
	_, err := s.client.PutObject(input)
	if err != nil {
		return "", err
	}
	return s.domain + objectKey, nil
}

func (s *huaweiobsServant) DeleteObject(objectKey string) error {
	_, err := s.client.DeleteObject(&obs.DeleteObjectInput{
		Bucket: s.bucket,
		Key:    objectKey,
	})
	return err
}

func (s *huaweiobsServant) DeleteObjects(objectKeys []string) error {
	input := &obs.DeleteObjectsInput{
		Bucket:  s.bucket,
		Objects: make([]obs.ObjectToDelete, 0, len(objectKeys)),
	}

	for _, key := range objectKeys {
		input.Objects = append(input.Objects, obs.ObjectToDelete{Key: key})
	}

	_, err := s.client.DeleteObjects(input)
	return err
}

func (s *huaweiobsServant) IsObjectExist(objectKey string) (bool, error) {
	input := &obs.GetObjectMetadataInput{
		Bucket: s.bucket,
		Key:    objectKey,
	}
	if _, err := s.client.GetObjectMetadata(input); err != nil {
		return false, err
	}
	return true, nil
}

func (s *huaweiobsServant) SignURL(objectKey string, expiredInSec int64) (string, error) {
	input := &obs.CreateSignedUrlInput{
		Method:  obs.HttpMethodGet,
		Bucket:  s.bucket,
		Key:     objectKey,
		Expires: int(expiredInSec),
	}
	out, err := s.client.CreateSignedUrl(input)
	if err != nil {
		logrus.Errorf("huaweiobsServant.SignURL err: %v", err)
		return "", err
	}

	ur, err := url.Parse(out.SignedUrl)
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

func (s *huaweiobsServant) ObjectURL(objetKey string) string {
	return s.domain + objetKey
}

func (s *huaweiobsServant) ObjectKey(objectUrl string) string {
	return strings.Replace(objectUrl, s.domain, "", -1)
}
