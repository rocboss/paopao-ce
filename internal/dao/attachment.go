package dao

import "github.com/rocboss/paopao-ce/internal/model"

func (d *dataServant) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	return attachment.Create(d.engine)
}
