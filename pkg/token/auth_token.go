// 系統登入驗證用
package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthTokenClaims struct {
	jwt.RegisteredClaims
}

// 是否有效
func (c *AuthTokenClaims) IsValid(nowTime time.Time) bool {
	// 檢查是否到期
	exp, _ := c.GetExpirationTime()
	if exp.Compare(nowTime) < 0 {
		return false
	}

	// 取生效時間
	validTime, _ := c.GetNotBefore()

	// 檢查生效時間是否早於現在時間
	return validTime.Compare(nowTime) <= 0
}

func (c *AuthTokenClaims) UID() string {
	return c.ID
}

func (c *AuthTokenClaims) Username() string {
	return c.Subject
}

// 解析Auth Token
func ParseAuthToken(tokenString string, secretKey []byte) (*AuthTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, err
	}

	return token.Claims.(*AuthTokenClaims), nil
}

// 產生Auth Token
func NewAuthToken(uid, username string, secretKey []byte, expirationTime time.Time, nowTime time.Time) (tokenString string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			Subject:   username,
			ID:        uid,
		},
	})

	tokenString, _ = token.SignedString(secretKey)

	return
}
