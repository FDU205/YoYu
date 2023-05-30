package posts

import (
	"YOYU/backend/utils"

	"github.com/gin-gonic/gin"
)

// validator 在验证用户后把对应的数据模型填好
type PostValidator struct {
	MessageBoxID uint   `json:"message_box_id" binding:"required"`
	PosterID     uint   `json:"userID" binding:"-"`
	Content      string `json:"content" binding:"required"`
	Visibility   uint   `json:"visibility" binding:"required,min=1,max=2"`
	PostModel    Post   `json:"-"`
}

// 将对应数据类型填好
func (PV *PostValidator) Bind(c *gin.Context) error {
	err := utils.Bind(c, PV)
	if err != nil {
		return err
	}
	PV.PostModel.MessageBoxID = PV.MessageBoxID
	PV.PostModel.PosterID = c.MustGet("userID").(uint)
	PV.PostModel.Content = PV.Content
	PV.PostModel.Visibility = PV.Visibility
	return nil
}

// 创建一个新的validator
func NewPostValidator() PostValidator {
	postValidator := PostValidator{}
	return postValidator
}

// validator 在验证用户后把对应的数据模型填好
type ChannelValidator struct {
	OwnerID      uint    `json:"userID" binding:"-"`
	PostID       uint    `json:"postID" binding:"required"`
	Content      string  `json:"content" binding:"required"`
	ChannelModel Channel `json:"-"`
}

// 将对应数据类型填好
func (CV *ChannelValidator) Bind(c *gin.Context) error {
	err := utils.Bind(c, CV)
	if err != nil {
		return err
	}
	CV.OwnerID = c.MustGet("userID").(uint)
	CV.ChannelModel.PostID = CV.PostID
	CV.ChannelModel.Content = CV.Content
	return nil
}

// 创建一个新的validator
func NewChannelValidator() ChannelValidator {
	channelValidator := ChannelValidator{}
	return channelValidator
}
