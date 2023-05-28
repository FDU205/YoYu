package utils

import (
	"YOYU/backend/common"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Id uint
	jwt.RegisteredClaims
}

// 根据id生成token
func GenToken(id uint) string {

	claim := MyClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(common.EXP)),  //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(common.TOKEN_KEY))
	return token
}

// 解析token， 弃用
func ParseToken(token string) (*MyClaims, error) {
	parser, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.TOKEN_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if !parser.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := parser.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}

func Token2Claims(token *jwt.Token) (*MyClaims, error) {
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
