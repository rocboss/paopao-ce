// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core/ms"
)

type Claims struct {
	UID      int64  `json:"uid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(conf.JWTSetting.Secret)
}

func GenerateToken(User *ms.User) (string, error) {
	expireTime := time.Now().Add(conf.JWTSetting.Expire)
	claims := Claims{
		UID:      User.ID,
		Username: User.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    conf.JWTSetting.Issuer + ":" + User.Salt,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (res *Claims, err error) {
	var tokenClaims *jwt.Token
	tokenClaims, err = jwt.ParseWithClaims(token, &Claims{}, func(_ *jwt.Token) (any, error) {
		return GetJWTSecret(), nil
	})
	if err == nil && tokenClaims != nil && tokenClaims.Valid {
		res, _ = tokenClaims.Claims.(*Claims)
	} else {
		err = jwt.ErrTokenNotValidYet
	}
	return
}
