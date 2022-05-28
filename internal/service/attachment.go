package service

import "github.com/rocboss/paopao-ce/internal/model"

func (svc *Service) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	return svc.dao.CreateAttachment(attachment)
}
