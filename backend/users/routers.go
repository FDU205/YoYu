package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 将Users模块的功能注册进框架
func UsersRegister(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
}

// 注册
func Register(c *gin.Context) {
	userValidator := NewRegisterValidator()
	if err := userValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"err_msg": err.Error()})
		return
	}

	if err := UserRegister(&userValidator.UserModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.Set("userModel", userValidator.UserModel)
	serializer := RegisterSerializer{c}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()})
}

// 登陆
func Login(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}
	userModel, err := UserLogin(loginValidator.UserModel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.Set("id", userModel.ID)
	serializer := LoginSerializer{c}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()})
}
