package users

import (
	"YOYU/backend/utils"

	"github.com/gin-gonic/gin"
)

type RegisterSerializer struct {
	c *gin.Context
}

type RegisterResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (r *RegisterSerializer) Response() RegisterResponse {
	myUserModel := r.c.MustGet("userModel").(User)
	user := RegisterResponse{
		Username: myUserModel.Username,
		Token:    utils.GenToken(myUserModel.ID),
	}
	return user
}

type LoginSerializer struct {
	c *gin.Context
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (LS *LoginSerializer) Response() LoginResponse {
	id := LS.c.MustGet("id").(uint)
	user := LoginResponse{
		Token: utils.GenToken(id),
	}
	return user
}
