// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/bitbus/sqlx"
	"github.com/bitbus/sqlx/db"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/auto/cc"
)

var (
	_ core.SecurityService = (*securitySrv)(nil)
)

type securitySrv struct {
	*sqlxSrv
	q           *cc.Security
	rand        *rand.Rand
	phoneVerify core.PhoneVerifyService
}

// GetLatestPhoneCaptcha 获取最新短信验证码
func (s *securitySrv) GetLatestPhoneCaptcha(phone string) (*ms.Captcha, error) {
	return db.Get[ms.Captcha](s.q.GetLatestPhoneCaptcha, phone)
}

// UsePhoneCaptcha 更新短信验证码
func (s *securitySrv) UsePhoneCaptcha(r *ms.Captcha) error {
	_, err := s.q.UsePhoneCaptcha.Exec(time.Now().Unix(), r.ID)
	return err
}

// SendPhoneCaptcha 发送短信验证码
func (s *securitySrv) SendPhoneCaptcha(phone string) error {
	expire := time.Duration(5)
	// 发送验证码
	captcha := strconv.Itoa(s.rand.Intn(900000) + 100000)
	if err := s.phoneVerify.SendPhoneCaptcha(phone, captcha, expire); err != nil {
		return err
	}
	// 写入表
	phoneCaptcha := &dbr.Captcha{
		Model: &ms.Model{
			CreatedOn: time.Now().Unix(),
		},
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
		q:           ccBuild(db, cc.BuildSecurity),
		rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
		phoneVerify: phoneVerify,
	}
}
