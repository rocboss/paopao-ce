package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

// StorageService storage service interface that implement base AliOSS、MINIO or other
type StorageService interface {
	CreateAttachment(attachment *model.Attachment) (*model.Attachment, error)
}
