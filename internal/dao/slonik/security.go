// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package slonik

import (
	"github.com/jackc/pgx/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/debug"
)

var (
	_ core.SecurityService = (*securityServant)(nil)
)

type securityServant struct {
	*pgxServant
}

// GetLatestPhoneCaptcha 获取最新短信验证码
func (s *securityServant) GetLatestPhoneCaptcha(phone string) (*core.Captcha, error) {
	// TODO
	debug.NotImplemented()
	return nil, nil
}

// UsePhoneCaptcha 更新短信验证码
func (s *securityServant) UsePhoneCaptcha(captcha *core.Captcha) error {
	// TODO
	debug.NotImplemented()
	return nil
}

// SendPhoneCaptcha 发送短信验证码
func (s *securityServant) SendPhoneCaptcha(phone string) error {
	// TODO
	debug.NotImplemented()
	return nil
}

func newSecurityService(db *pgx.Conn, phoneVerify core.PhoneVerifyService) core.SecurityService {
	return &securityServant{
		pgxServant: newPgxServant(db),
	}
}
