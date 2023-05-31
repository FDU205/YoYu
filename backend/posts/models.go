package posts

import (
	"YOYU/backend/database"
	"YOYU/backend/messagebox"
	"YOYU/backend/users"
)

type Post struct {
	ID           uint                  `gorm:"primary_key" json:"id"`
	User         users.User            `gorm:"ForeignKey:PosterID" json:"-"`
	PosterID     uint                  `gorm:"column:poster_id; not null" json:"poster_id"`
	MessageBox   messagebox.MessageBox `gorm:"ForeignKey:MessageBoxID" json:"-"`
	MessageBoxID uint                  `gorm:"column:message_box_id; not null" json:"message_box_id"`
	Content      string                `gorm:"column:content; not null" json:"content"`
	Visibility   uint                  `gorm:"column:visibility; not null" json:"visibility"`
}

type Channel struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Post    Post   `gorm:"ForeignKey:PostID" json:"-"`
	PostID  uint   `gorm:"column:post_id; not null" json:"post_id"`
	Content string `gorm:"column:content; not null" json:"content"`
	Type    uint   `gorm:"column:type; not null" json:"type"`
}

// 获取回复
func GetChannels(postID uint) ([]Channel, error) {
	db := database.GetDB()
	var channels []Channel
	err := db.Where("post_id = ?", postID).Order("id").Find(&channels).Error
	return channels, err
}

// 创建回复
func CreateChannel(channelModel *Channel) error {
	db := database.GetDB()
	err := db.Save(channelModel).Error
	return err
}

// 创建帖子
func CreatePost(postModel *Post) error {
	db := database.GetDB()
	err := db.Save(postModel).Error
	return err
}

// 根据帖子ID获取帖子
func GetPostByID(postID uint) (Post, error) {
	db := database.GetDB()
	var post Post
	err := db.Where("id = ?", postID).First(&post).Error
	return post, err
}

// 根据用户ID获取帖子
func GetPostByUserID(userID uint, offset int, limit int) ([]Post, error) {
	db := database.GetDB()
	var posts []Post
	err := db.Where("poster_id = ?", userID).Order("id desc").Limit(limit).Offset(offset).Find(&posts).Error
	return posts, err
}

// 查询帖子
func SearchPost(message_box_id uint, offset int, limit int) ([]Post, error) {
	db := database.GetDB()
	var posts []Post
	err := db.Where("message_box_id = ?", message_box_id).Order("id desc").Limit(limit).Offset(offset).Find(&posts).Error
	return posts, err
}

// 根据帖子ID删除帖子
func DeletePostByID(postID uint, posterID uint) error {
	db := database.GetDB()
	err := db.Where("id = ? AND poster_id = ?", postID, posterID).Delete(Post{}).Error
	return err
}
