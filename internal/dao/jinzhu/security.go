// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

var (
	_ core.SecurityService = (*securityServant)(nil)
)

type securityServant struct {
	db          *gorm.DB
	phoneVerify core.PhoneVerifyService
}

func newSecurityService(db *gorm.DB, phoneVerify core.PhoneVerifyService) core.SecurityService {
	return &securityServant{
		db:          db,
		phoneVerify: phoneVerify,
	}
}

// GetLatestPhoneCaptcha 获取最新短信验证码
func (s *securityServant) GetLatestPhoneCaptcha(phone string) (*core.Captcha, error) {
	return (&dbr.Captcha{
		Phone: phone,
	}).Get(s.db)
}

// UsePhoneCaptcha 更新短信验证码
func (s *securityServant) UsePhoneCaptcha(captcha *core.Captcha) error {
	captcha.UseTimes++
	return captcha.Update(s.db)
}

// SendPhoneCaptcha 发送短信验证码
func (s *securityServant) SendPhoneCaptcha(phone string) error {
	expire := time.Duration(5)

	// 发送验证码
	rand.Seed(time.Now().UnixNano())
	captcha := strconv.Itoa(rand.Intn(900000) + 100000)
	if err := s.phoneVerify.SendPhoneCaptcha(phone, captcha, expire); err != nil {
		return err
	}

	// 写入表
	captchaModel := &dbr.Captcha{
		Phone:     phone,
		Captcha:   captcha,
		ExpiredOn: time.Now().Add(expire * time.Minute).Unix(),
	}
	captchaModel.Create(s.db)
	return nil
}
