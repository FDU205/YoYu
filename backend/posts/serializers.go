package posts

import (
	"YOYU/backend/users"

	"github.com/gin-gonic/gin"
)

type SearchSerializer struct {
	Posts []Post `json:"posts"`
}

type GetSerializer struct {
	c *gin.Context
}

type GetResponse struct {
	ID           uint      `json:"id"`
	PosterID     uint      `json:"poster_id"`
	PosterName   string    `json:"poster_name"`
	Content      string    `json:"content"`
	Visibility   uint      `json:"visibility"`
	MessageBoxID uint      `json:"message_box_id"`
	Threads      []Channel `json:"threads"`
	Channels     []string  `json:"channels"`
}

func (r *GetSerializer) Response() GetResponse {
	mypostModel := r.c.MustGet("postModel").(Post)
	myThreads := r.c.MustGet("channels").([]Channel)

	var postname string
	if mypostModel.Visibility == 2 {
		postname = "匿名用户"
	} else {
		user, _ := users.GetUser(&users.User{ID: mypostModel.PosterID})
		postname = user.Username
	}

	ret := GetResponse{
		ID:           mypostModel.ID,
		PosterID:     mypostModel.PosterID,
		PosterName:   postname,
		Content:      mypostModel.Content,
		Visibility:   mypostModel.Visibility,
		MessageBoxID: mypostModel.MessageBoxID,
		Threads:      myThreads,
		Channels:     []string{},
	}
	return ret
}
