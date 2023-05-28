package users

import (
	"YOYU/backend/common"
	"YOYU/backend/database"
	"YOYU/backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

// 从token中剥去bearer前缀
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	if len(tok) > 6 && strings.ToUpper(tok[0:6]) == "bearer " {
		return tok[7:], nil
	}
	return tok, nil
}

// 从Authorization头中获取token
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// 鉴权之后将对应用户填入context
func UpdateContextUserModel(c *gin.Context, userId uint) {
	var userModel User
	if userId != 0 {
		db := database.GetDB()
		db.First(&userModel, userId)
	}
	c.Set("userModel", userModel)
}

// 鉴权中间件
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(common.TOKEN_KEY), nil
		})
		if err != nil || token.Valid {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		claims, err := utils.Token2Claims(token)
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		} else {
			userId := claims.Id
			UpdateContextUserModel(c, userId)
		}
	}
}
