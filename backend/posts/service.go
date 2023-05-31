package posts

import (
	"YOYU/backend/messagebox"
	"YOYU/backend/users"
	"errors"
)

// 创建帖子
func PostCreate(postModel *Post) error {
	err := CreatePost(postModel)
	return err
}

// 根据ID获取帖子
func PostGetByID(postID uint) ([]Channel, Post, error) {
	post, err := GetPostByID(postID)
	if err != nil {
		return nil, post, err
	}
	channels, err := GetChannels(postID)
	return channels, post, err
}

// 搜索帖子
func PostSearch(message_box_id uint, pageNum int, pageSize int) ([]Post, error) {
	offset := (pageNum - 1) * pageSize
	limit := pageSize
	post, err := SearchPost(message_box_id, offset, limit)
	return post, err
}

// 获取我的帖子
func MyPostGet(userID uint, pageNum int, pageSize int) ([]Post, error) {
	offset := (pageNum - 1) * pageSize
	limit := pageSize
	post, err := GetPostByUserID(userID, offset, limit)
	return post, err
}

// 根据ID删除帖子
func PostDeleteByID(postID uint, ownerID uint) error {
	err := DeletePostByID(postID, ownerID)
	return err
}

// 创建回复
func ChannelCreate(owner_id uint, channelModel *Channel) error {
	post, err := GetPostByID(channelModel.PostID)
	if err != nil {
		return err
	}
	// 只能追问自己的帖子
	if channelModel.Type == 1 && owner_id != post.PosterID {
		return errors.New("无权限追问")
	}
	messagebox, err := messagebox.MessageBoxGetByID(post.MessageBoxID)
	if err != nil {
		return errors.New("错误")
	}
	user, err := users.GetUser(&users.User{ID: messagebox.OwnerID})
	if err != nil {
		return errors.New("错误")
	}

	// 只能回答自己提问箱里的问题
	if channelModel.Type == 2 && user.ID != owner_id {
		return errors.New("无权限回复")
	}
	err = CreateChannel(channelModel)
	return err
}
