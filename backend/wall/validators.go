package wall

import (
	"YOYU/backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// validator 在验证用户后把对应的数据模型填好
type WallValidator struct {
	PosterID   uint   `json:"userID" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Visibility uint   `json:"visibility" binding:"required,min=0,max=1"`
	WallModel  Wall   `json:"-"`
}

// 将对应数据类型填好
func (WV *WallValidator) Bind(c *gin.Context) error {
	err := utils.Bind(c, WV)
	if err != nil {
		return err
	}
	WV.WallModel.PosterID = WV.PosterID
	WV.WallModel.Content = WV.Content
	WV.WallModel.Visibility = WV.Visibility
	WV.WallModel.Date = time.Now().Format("2006-01-02")
	return nil
}

// 创建一个新的validator
func NewWallValidator() WallValidator {
	wallValidator := WallValidator{}
	return wallValidator
}
