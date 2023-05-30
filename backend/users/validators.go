package users

import (
	"YOYU/backend/utils"

	"github.com/gin-gonic/gin"
)

// validator 在验证用户后把对应的数据模型填好
type RegisterValidator struct {
	Username  string `json:"username" binding:"required,min=1,max=32"`
	Password  string `json:"password" binding:"required,min=1,max=255"`
	UserModel User   `json:"-"`
}

// 将对应数据类型填好
func (RV *RegisterValidator) Bind(c *gin.Context) error {
	err := utils.Bind(c, RV)
	if err != nil {
		return err
	}
	RV.UserModel.Username = RV.Username
	RV.UserModel.SetPassword(RV.Password)
	return nil
}

// 创建一个新的Uservalidator
func NewRegisterValidator() RegisterValidator {
	registerValidator := RegisterValidator{}
	return registerValidator
}

// 登陆时用的Validator
type LoginValidator struct {
	Username  string `json:"username" binding:"required,alphanum,min=1,max=255"`
	Password  string `json:"password" binding:"required,min=1,max=255"`
	UserModel User   `json:"-"`
}

func (LV *LoginValidator) Bind(c *gin.Context) error {
	err := utils.Bind(c, LV)
	if err != nil {
		return err
	}

	LV.UserModel.Username = LV.Username
	LV.UserModel.Password = LV.Password
	return nil
}

// 创建一个新的Loginvalidator
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
