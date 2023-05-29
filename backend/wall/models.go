package wall

import (
	"YOYU/backend/database"
)

type Wall struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Poster_id  uint   `gorm:"column:poster_id; not null" json:"poster_id"`
	Content    string `gorm:"column:body; not null" json:"content"`
	Visibility uint   `gorm:"column:visibility; not null" json:"visibility"`
	Date       string `gorm:"column:date; not null" json:"-"`
}

// 创建表白信息
func CreateWall(data interface{}) error {
	db := database.GetDB()
	err := db.Save(data).Error
	return err
}

// 返回今日所有表白信息
func GetWall(date string, pageNum int, pageSize int) ([]Wall, error) {
	offset := (pageNum - 1) * pageSize
	limit := pageSize

	db := database.GetDB()
	var wall []Wall

	err := db.Where("Date = ?", date).Order("id desc").Limit(limit).Offset(offset).Find(&wall).Error
	return wall, err
}
