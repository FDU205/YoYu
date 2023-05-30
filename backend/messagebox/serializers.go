package messagebox

import (
	"github.com/gin-gonic/gin"
)

type SearchSerializer struct {
	MessageBoxes []MessageBox `json:"messageBoxes"`
}

type GetSerializer struct {
	c *gin.Context
}

type GetResponse struct {
	ID      uint   `json:"id"`
	OwnerID uint   `json:"owner_id"`
	Title   string `json:"title"`
	Posts   []uint `json:"posts"`
}

func (r *GetSerializer) Response() GetResponse {
	myMessageBox := r.c.MustGet("messageBoxModel").(MessageBox)
	myPosts := r.c.MustGet("posts").([]uint)
	ret := GetResponse{
		ID:      myMessageBox.ID,
		OwnerID: myMessageBox.OwnerID,
		Title:   myMessageBox.Title,
		Posts:   myPosts,
	}
	return ret
}
