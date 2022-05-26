package dao

import "github.com/rocboss/paopao-api/internal/model"

func (d *Dao) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	return attachment.Create(d.engine)
}
