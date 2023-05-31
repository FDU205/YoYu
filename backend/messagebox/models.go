package messagebox

import (
	"YOYU/backend/database"
	"YOYU/backend/users"

	"gorm.io/gorm"
)

type MessageBox struct {
	gorm.Model `json:"-"`
	ID         uint       `gorm:"primary_key" json:"id"`
	User       users.User `gorm:"ForeignKey:OwnerID" json:"-"`
	OwnerID    uint       `gorm:"column:owner_id; not null" json:"owner_id"`
	Title      string     `gorm:"column:title; not null" json:"title"`
}

// 创建提问箱
func CreateMessageBox(messageBoxModel *MessageBox) error {
	db := database.GetDB()
	err := db.Save(messageBoxModel).Error
	return err
}

// 根据提问箱ID获取提问箱
func GetMessageBoxByID(messageBoxID uint) (MessageBox, error) {
	db := database.GetDB()
	var messageBox MessageBox
	err := db.Where("id = ?", messageBoxID).First(&messageBox).Error
	return messageBox, err
}

// 查询提问箱
func SearchMessageBox(title string, ownerID uint, offset int, limit int) ([]MessageBox, error) {
	db := database.GetDB()
	var messageBox []MessageBox
	query := db

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	if ownerID != 0 {
		query = query.Where("owner_id = ?", ownerID)
	}

	err := query.Order("updated_at desc").Limit(limit).Offset(offset).Find(&messageBox).Error
	return messageBox, err
}

// 根据提问箱ID删除提问箱
func DeleteMessageBoxByID(messageBoxID uint, ownerID uint) error {
	db := database.GetDB()
	err := db.Where("id = ? AND owner_id = ?", messageBoxID, ownerID).Delete(&MessageBox{}).Error
	return err
}

// 根据提问箱ID更新提问箱
func UpdateMessageBoxByID(data interface{}) error {
	db := database.GetDB()
	err := db.Save(data).Error
	return err
}
