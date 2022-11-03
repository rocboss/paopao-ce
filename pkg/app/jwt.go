package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/model"
)

type Claims struct {
	UID      int64  `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(conf.JWTSetting.Secret)
}

func GenerateToken(User *model.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(conf.JWTSetting.Expire)
	claims := Claims{
		UID:      User.ID,
		Username: User.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    conf.JWTSetting.Issuer + ":" + User.Salt,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
