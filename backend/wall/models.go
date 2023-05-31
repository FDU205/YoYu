package wall

import (
	"YOYU/backend/database"
	"YOYU/backend/users"
)

type Wall struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	User       users.User `gorm:"ForeignKey:PosterID" json:"-"`
	PosterID   uint       `gorm:"column:poster_id; not null" json:"poster_id"`
	Content    string     `gorm:"column:content; not null" json:"content"`
	Visibility uint       `gorm:"column:visibility; not null" json:"visibility"`
	Date       string     `gorm:"column:date; not null" json:"-"`
}

// 创建表白信息
func CreateWall(data interface{}) error {
	db := database.GetDB()
	err := db.Save(data).Error
	return err
}

// 返回今日所有表白信息
func GetWall(date string, offset int, limit int) ([]Wall, error) {

	db := database.GetDB()
	var wall []Wall

	err := db.Where("Date = ?", date).Order("id desc").Limit(limit).Offset(offset).Find(&wall).Error
	return wall, err
}

// 返回我的表白信息
func GetWallbyID(id uint, offset int, limit int) ([]Wall, error) {

	db := database.GetDB()
	var wall []Wall

	err := db.Where("poster_id = ?", id).Order("id desc").Limit(limit).Offset(offset).Find(&wall).Error
	return wall, err
}
