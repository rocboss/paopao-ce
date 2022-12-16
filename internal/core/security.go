// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"time"

	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

type (
	Captcha = dbr.Captcha
)

// SecurityService 安全相关服务
type SecurityService interface {
	GetLatestPhoneCaptcha(phone string) (*Captcha, error)
	UsePhoneCaptcha(captcha *Captcha) error
	SendPhoneCaptcha(phone string) error
}

// AttachmentCheckService 附件检测服务
type AttachmentCheckService interface {
	CheckAttachment(uri string) error
}

// PhoneVerifyService 手机验证服务
type PhoneVerifyService interface {
	SendPhoneCaptcha(phone string, captcha string, expire time.Duration) error
}
