// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid/v5"
	"github.com/rocboss/paopao-ce/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type PasswordProvider interface {
	Generate(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type bcryptPasswordProvider struct {
	cost int
}

type md5PasswordProvider struct{}

func (p *bcryptPasswordProvider) Generate(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	return utils.String(hashedPassword), err
}

func (p *bcryptPasswordProvider) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (p md5PasswordProvider) Generate(password string) (string, error) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	return utils.EncodeMD5(utils.EncodeMD5(password)+salt) + ":" + salt, nil
}

func (p md5PasswordProvider) Compare(hashedPassword, password string) error {
	passwordSalt := strings.Split(string(hashedPassword), ":")
	if len(passwordSalt) != 2 {
		return errors.New("invalid hashed password")
	}
	if strings.Compare(passwordSalt[0], utils.EncodeMD5(utils.EncodeMD5(password)+passwordSalt[1])) != 0 {
		return errors.New("invalid password")
	}
	return nil
}

func NewBcryptPasswordProvider(cost int) PasswordProvider {
	return &bcryptPasswordProvider{
		cost: cost,
	}
}

func NewMd5PasswordProvider() PasswordProvider {
	return md5PasswordProvider{}
}
