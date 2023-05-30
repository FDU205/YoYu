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

	claims := make(jwt.MapClaims) //数据仓声明
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(common.EXP))
	claims["userID"] = id

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(common.TOKEN_KEY))
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

// 已弃用
// 将token转换为Claims
func Token2Claims(token *jwt.Token) (*MyClaims, error) {
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
