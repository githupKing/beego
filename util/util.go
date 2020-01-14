/*
* @Author: wangyongping
* @Date:   2020-01-12 22:03:16
* @Last Modified by:   Administrator
* @Last Modified time: 2020-01-14 00:27:43
 */
package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claims custom token
type Claims struct {
	Uid int64 `json:"u_id"` // 用户id
	jwt.StandardClaims
}

// CreateToken create token
func CreateToken(claims *Claims) (signedToken string, success bool) {
	claims.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return
	}
	success = true
	return
}

// ValidateToken validate token
func ValidateToken(signedToken string) (claims *Claims, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		success = true
		return
	}

	return
}
