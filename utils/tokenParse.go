package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

var stSigningKey = []byte(viper.GetString("jwt.signingKey"))

type JwtCustClaims struct {
	UserID   string
	ID       string
	UserName string
	Email    string
	Phone    string
	WxID     string
	jwt.RegisteredClaims
}

func GenerateToken(userid string) (string, error) { //生成Token
	iJwtCustClaims := JwtCustClaims{
		UserID: userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), //过期时间设置为三十分钟
			IssuedAt:  jwt.NewNumericDate(time.Now()),                       // 签发时间
			Subject:   "Token",                                              //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	return token.SignedString(stSigningKey)
}

func ParseToken(tokenStr string) (*JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return stSigningKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("Invalid Token")
	}
	return &iJwtCustClaims, nil
}

// 判断Token是否正确

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}
