package dao

import (
	"fmt"
	"strings"

	"github.com/rocboss/paopao-ce/internal/model"
)

func (d *dataServant) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	return attachment.Create(d.engine)
}

func (s *attachmentCheckServant) CheckAttachment(uri string) error {
	if strings.Index(uri, s.domain) != 0 {
		return fmt.Errorf("附件非本站资源")
	}
	return nil
}
