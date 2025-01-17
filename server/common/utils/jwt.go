package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	ID    int64
	State string `json:"state"`
	jwt.StandardClaims
}

// GetToken 获取accessToken和refreshToken
func GetToken(id int64, state, a, r string) (string, string) {
	// accessToken 的数据
	accessSecret := []byte(a)
	refreshSecret := []byte(r)
	aT := MyClaims{
		id,
		state,
		jwt.StandardClaims{
			Issuer:    "AR",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
		},
	}
	// refreshToken 的数据

	rT := MyClaims{
		id,
		state,
		jwt.StandardClaims{
			Issuer:    "RT",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aT)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rT)
	accessTokenSigned, err := accessToken.SignedString(accessSecret)
	if err != nil {
		fmt.Println("获取Token失败，Secret错误")
		return "", ""
	}
	refreshTokenSigned, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		fmt.Println("获取Token失败，Secret错误")
		return "", ""
	}
	return accessTokenSigned, refreshTokenSigned
}

func ParseToken(accessTokenString, refreshTokenString, a, r string) (*MyClaims, bool, error) {
	accessSecret := []byte(a)
	refreshSecret := []byte(r)
	accessToken, err := jwt.ParseWithClaims(accessTokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})
	if claims, ok := accessToken.Claims.(*MyClaims); ok && accessToken.Valid {
		return claims, false, nil
	}

	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})
	if err != nil {
		return nil, false, err
	}
	if claims, ok := refreshToken.Claims.(*MyClaims); ok && refreshToken.Valid {
		return claims, true, nil
	}

	return nil, false, errors.New("invalid token")
}
