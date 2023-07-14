// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
)

var (
	_ core.SecurityService = (*securitySrv)(nil)
)

type securitySrv struct {
	*sqlxSrv
	q           *cc.Security
	phoneVerify core.PhoneVerifyService
}

// GetLatestPhoneCaptcha 获取最新短信验证码
func (s *securitySrv) GetLatestPhoneCaptcha(phone string) (*ms.Captcha, error) {
	res := &ms.Captcha{}
	err := s.q.GetLatestPhoneCaptcha.Get(res, phone)
	return res, err
}

// UsePhoneCaptcha 更新短信验证码
func (s *securitySrv) UsePhoneCaptcha(r *ms.Captcha) error {
	r.UseTimes++
	_, err := s.q.UsePhoneCaptcha.Exec(r)
	return err
}

// SendPhoneCaptcha 发送短信验证码
func (s *securitySrv) SendPhoneCaptcha(phone string) error {
	expire := time.Duration(5)
	// 发送验证码
	rand.Seed(time.Now().UnixNano())
	captcha := strconv.Itoa(rand.Intn(900000) + 100000)
	if err := s.phoneVerify.SendPhoneCaptcha(phone, captcha, expire); err != nil {
		return err
	}
	// 写入表
	phoneCaptcha := &dbr.Captcha{
		Phone:     phone,
		Captcha:   captcha,
		ExpiredOn: time.Now().Add(expire * time.Minute).Unix(),
	}
	_, err := s.q.CreatePhoneCaptcha.Exec(phoneCaptcha)
	return err
}

func newSecurityService(db *sqlx.DB, phoneVerify core.PhoneVerifyService) core.SecurityService {
	return &securitySrv{
		sqlxSrv:     newSqlxSrv(db),
		q:           mustBuild(db, cc.BuildSecurity),
		phoneVerify: phoneVerify,
	}
}
