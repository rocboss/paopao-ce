// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"github.com/jmoiron/sqlx"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/sakila/yesql/cc"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.SecurityService = (*securitySrv)(nil)
)

type securitySrv struct {
	*sqlxSrv
	q *cc.Security
}

// GetLatestPhoneCaptcha 获取最新短信验证码
func (s *securitySrv) GetLatestPhoneCaptcha(phone string) (*core.Captcha, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

// UsePhoneCaptcha 更新短信验证码
func (s *securitySrv) UsePhoneCaptcha(captcha *core.Captcha) error {
	// TODO
	debug.NotImplemented()
	return nil
}

// SendPhoneCaptcha 发送短信验证码
func (s *securitySrv) SendPhoneCaptcha(phone string) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func newSecurityService(db *sqlx.DB, phoneVerify core.PhoneVerifyService) core.SecurityService {
	return &securitySrv{
		sqlxSrv: newSqlxSrv(db),
		q:       mustBuild(db, cc.BuildSecurity),
	}
}
