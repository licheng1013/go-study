package common

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 定义一个密钥，用于签名和验证JWT
var secretKey = []byte("uIiwiZXhwIjoxNjc4MDkyMTg2LCJ")

func CreateToken(ID string) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute) //这是设置过期时间,10分钟
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		//Issuer:    "WebToken", //可空
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ID:       ID,
		//Subject:   "MyInfo",  //可空
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析并验证JWT，并检查是否过期
func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("验证失败")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 如果token有效，将其断言为自定义的声明结构体，并返回其中的信息
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
