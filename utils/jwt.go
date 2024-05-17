package utils

import (
	"SCIProj/global"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Uid string
	jwt.StandardClaims
}

// 生成Token

func Award(uid string) (string, error) {
	// 过期时间 默认3天
	expireTime := time.Now().Add(3 * 24 * time.Hour)
	claims := &Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(global.JWTKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// 解析token

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return global.JWTKey, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, claims, err
}
