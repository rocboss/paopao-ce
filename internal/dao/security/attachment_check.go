// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	"fmt"
	"strings"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
)

var (
	_ core.AttachmentCheckService = (*attachmentCheckServant)(nil)
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
