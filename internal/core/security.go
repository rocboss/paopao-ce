package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

// SecurityService 安全相关服务
type SecurityService interface {
	GetLatestPhoneCaptcha(phone string) (*model.Captcha, error)
	UsePhoneCaptcha(captcha *model.Captcha) error
	SendPhoneCaptcha(phone string) error
}

// 附件检测服务
type AttachmentCheckService interface {
	CheckAttachment(uri string) error
}
