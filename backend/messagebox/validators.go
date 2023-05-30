package messagebox

import (
	"YOYU/backend/utils"

	"github.com/gin-gonic/gin"
)

// validator 在验证用户后把对应的数据模型填好
type MessageBoxValidator struct {
	OwnerID         uint       `json:"userID" binding:"-"`
	Title           string     `json:"title" binding:"required"`
	MessageBoxModel MessageBox `json:"-"`
}

// 将对应数据类型填好
func (MV *MessageBoxValidator) Bind(c *gin.Context) error {
	err := utils.Bind(c, MV)
	if err != nil {
		return err
	}
	MV.MessageBoxModel.OwnerID = c.MustGet("userID").(uint)
	MV.MessageBoxModel.Title = MV.Title
	return nil
}

// 创建一个新的validator
func NewMessageBoxValidator() MessageBoxValidator {
	MessageBoxValidator := MessageBoxValidator{}
	return MessageBoxValidator
}
