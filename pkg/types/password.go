// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordGenerator interface {
	Generate([]byte) ([]byte, error)
}

func NewBcryptPasswordGenerator(cost int) PasswordGenerator {
	return &bcryptPasswordGenerator{
		cost: cost,
	}
}

type bcryptPasswordGenerator struct {
	cost int
}

func (g *bcryptPasswordGenerator) Generate(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, g.cost)
}
