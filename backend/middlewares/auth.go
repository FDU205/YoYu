package middlewares

import (
	"YOYU/backend/common"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

// AuthorizationHeaderExtractor extracts a bearer token from Authorization header
// Uses PostExtractionFilter to strip "Bearer " prefix from header
var myAuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

// 鉴权中间件
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, myAuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(common.TOKEN_KEY), nil
		})
		if err != nil || !token.Valid {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, errors.New("授权失败或授权已过期"))
			}
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Set("userID", uint(claims["userID"].(float64)))
		c.Next()
	}
}
