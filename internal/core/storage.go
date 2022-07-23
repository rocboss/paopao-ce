package core

import (
	"io"
)

// ObjectStorageService storage service interface that implement base AliOSS、MINIO or other
type ObjectStorageService interface {
	PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string) (string, error)
	DeleteObject(objectKey string) error
	DeleteObjects(objectKeys []string) error
	IsObjectExist(objectKey string) (bool, error)
	SignURL(objectKey string, expiredInSec int64) (string, error)
	ObjectURL(objetKey string) string
	ObjectKey(cUrl string) string
}
