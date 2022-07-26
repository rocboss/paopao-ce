package core

import (
	"io"
)

// ObjectStorageService storage service interface that implement base AliOSS、MINIO or other
type ObjectStorageService interface {
	OssCreateService
	OssDeleteService

	SignURL(objectKey string, expiredInSec int64) (string, error)
	ObjectURL(objetKey string) string
	ObjectKey(cUrl string) string
}

// OssCreateService Object Storage System Object Create service
type OssCreateService interface {
	PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, persistance bool) (string, error)
	PersistObject(objectKey string) error
}

// OssCreateService Object Storage System Object Delete service
type OssDeleteService interface {
	DeleteObject(objectKey string) error
	DeleteObjects(objectKeys []string) error
	IsObjectExist(objectKey string) (bool, error)
}
