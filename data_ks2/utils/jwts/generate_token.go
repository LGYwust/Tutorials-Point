package jwts

import (
	"data_ks2/global"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenToken(user JwyPayLoad) (string, error) {
	Secret = []byte(global.Config.Jwt.Secret)
	claims := jwt.MapClaims{
		"iss":     global.Config.Jwt.Iss,
		"exp":     time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Exp)).Unix(),
		"payload": user,
	}
	// 创建一个JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对Token进行签名并获得完整的编码后的字符串Token
	tokenString, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
