package wall

import (
	"YOYU/backend/users"

	"github.com/gin-gonic/gin"
)

type GetSerializer struct {
	c *gin.Context
}

type Getin struct {
	ID         uint   `json:"id"`
	PosterID   uint   `json:"poster_id"`
	PosterName string `json:"poster_name"`
	Content    string `json:"content"`
	Visibility uint   `json:"visibility"`
}

type GetResponse struct {
	Posts []Getin `json:"posts"`
}

func (r *GetSerializer) Response() GetResponse {
	myWall := r.c.MustGet("wallModel").([]Wall)
	var myPosts = []Getin{}
	for _, wall := range myWall {
		var username string
		if wall.Visibility == 2 {
			username = "匿名用户"
		} else {
			user, _ := users.GetUser(&users.User{ID: wall.PosterID})
			username = user.Username
		}
		myPosts = append(myPosts, Getin{
			ID:         wall.ID,
			PosterID:   wall.PosterID,
			PosterName: username,
			Content:    wall.Content,
			Visibility: wall.Visibility,
		})
	}
	ret := GetResponse{
		Posts: myPosts,
	}
	return ret
}
