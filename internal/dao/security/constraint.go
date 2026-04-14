//go:build constraint

package security

import (
	"github.com/rocboss/paopao-ce/internal/core"
)

var (
	_ core.AttachmentCheckService = (*attachmentCheckServant)(nil)
	_ core.PhoneVerifyService     = (*juheSmsServant)(nil)
)
