package security

import (
	"strings"

	"github.com/alimy/cfg"
	"github.com/rocboss/paopao-ce/internal/core"
)

func NewPhoneVerifyService() core.PhoneVerifyService {
	smsVendor, _ := cfg.Val("sms")
	switch strings.ToLower(smsVendor) {
	case "smsjuhe":
		return newJuheSmsServant()
	default:
		return newJuheSmsServant()
	}
}
