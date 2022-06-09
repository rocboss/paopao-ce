package dao

import (
	"io"
	"net/url"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/rocboss/paopao-ce/internal/conf"
)

func newAliossServent() *aliossServant {
	client, err := oss.New(conf.AliOSSSetting.Endpoint, conf.AliOSSSetting.AccessKeyID, conf.AliOSSSetting.AccessKeySecret)
	if err != nil {
		conf.Logger.Fatalf("alioss.New err: %v", err)
	}

	bucket, err := client.Bucket(conf.AliOSSSetting.Bucket)
	if err != nil {
		conf.Logger.Fatalf("client.Bucket err: %v", err)
	}

	return &aliossServant{
		bucket: bucket,
		domain: getOssDomain(),
	}
}

func (s *aliossServant) PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	err := s.bucket.PutObject(objectKey, reader, oss.ContentLength(objectSize), oss.ContentType(contentType))
	if err != nil {
		return "", err
	}
	return s.domain + objectKey, nil
}

func (s *aliossServant) SignURL(objectKey string, expiredInSec int64) (string, error) {
	signedURL, err := s.bucket.SignURL(objectKey, oss.HTTPGet, expiredInSec)
	if err != nil {
		conf.Logger.Errorf("client.SignURL err: %v", err)
		return "", err
	}

	ur, err := url.Parse(signedURL)
	if err != nil {
		conf.Logger.Errorf("url.Parse err: %v", err)
		return "", err
	}

	epath, err := url.PathUnescape(ur.Path)
	if err != nil {
		conf.Logger.Errorf("url.PathUnescape err: %v", err)
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
