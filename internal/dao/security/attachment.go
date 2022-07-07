package security

import (
	"fmt"
	"strings"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
)

type attachmentCheckServant struct {
	domain string
}

func (s *attachmentCheckServant) CheckAttachment(uri string) error {
	if strings.Index(uri, s.domain) != 0 {
		return fmt.Errorf("附件非本站资源")
	}
	return nil
}

func NewAttachmentCheckService() core.AttachmentCheckService {
	return &attachmentCheckServant{
		domain: conf.GetOssDomain(),
	}
}
