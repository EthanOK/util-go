package utils

import (
	"time"
	"util-go/config"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Name     string
	PassWord string
	jwt.RegisteredClaims
}

func GetAuthorization(name, password string) string {

	return "Bearer " + GenerateToken(name, password)
}

func GenerateToken(name, password string) string {
	claim := Claims{
		Name:     name,
		PassWord: password,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "GOD",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.JWT_EXPIRATION))), // 设置token过期时间为24小时
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, e := token.SignedString([]byte(config.JWT_SECRET))
	if e != nil {
		return ""
	}
	return t

}

func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
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
