package messagebox

import (
	"YOYU/backend/users"

	"github.com/gin-gonic/gin"
)

type SearchSerializer struct {
	c *gin.Context
}

type SearchIn struct {
	ID       uint   `json:"id"`
	OwnerID  uint   `json:"owner_id"`
	Title    string `json:"title"`
	Username string `json:"owner_name"`
}

type SearchResponse struct {
	MessageBoxes []SearchIn `json:"messageBoxes"`
}

func (r *SearchSerializer) Response() SearchResponse {
	myMessageBoxes := r.c.MustGet("messageBoxes").([]MessageBox)
	mySearchIn := []SearchIn{}
	for _, myMessageBox := range myMessageBoxes {
		userModel, _ := users.GetUser(&users.User{ID: myMessageBox.OwnerID})
		mySearchIn = append(mySearchIn, SearchIn{
			ID:       myMessageBox.ID,
			OwnerID:  myMessageBox.OwnerID,
			Title:    myMessageBox.Title,
			Username: userModel.Username,
		})
	}

	return SearchResponse{
		MessageBoxes: mySearchIn,
	}
}

type GetSerializer struct {
	c *gin.Context
}

type GetResponse struct {
	ID        uint   `json:"id"`
	OwnerID   uint   `json:"owner_id"`
	OwnerName string `json:"owner_name"`
	Title     string `json:"title"`
	Posts     []uint `json:"posts"`
}

func (r *GetSerializer) Response() GetResponse {
	myMessageBox := r.c.MustGet("messageBoxModel").(MessageBox)
	myPosts := r.c.MustGet("posts").([]uint)
	user, _ := users.GetUser(&users.User{ID: myMessageBox.OwnerID})
	ret := GetResponse{
		ID:        myMessageBox.ID,
		OwnerID:   myMessageBox.OwnerID,
		OwnerName: user.Username,
		Title:     myMessageBox.Title,
		Posts:     myPosts,
	}
	return ret
}
