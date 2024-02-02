// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"github.com/alimy/tryst/cfg"
	"github.com/rocboss/paopao-ce/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

func NewPasswordProvider() (provider auth.PasswordProvider) {
	cfg.On(cfg.Actions{
		"Md5AuthMethod": func() {
			provider = auth.NewMd5PasswordProvider()
		},
		"BcryptAuthMethod": func() {
			provider = auth.NewBcryptPasswordProvider(bcrypt.DefaultCost)
		},
	},
		func() {
			provider = auth.NewMd5PasswordProvider()
		},
	)
	return
}
