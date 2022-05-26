package service

import "github.com/rocboss/paopao-api/internal/model"

func (svc *Service) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	return svc.dao.CreateAttachment(attachment)
}
