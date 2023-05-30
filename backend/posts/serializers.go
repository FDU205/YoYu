package posts

import "github.com/gin-gonic/gin"

type SearchSerializer struct {
	Posts []Post `json:"posts"`
}

type GetSerializer struct {
	c *gin.Context
}

type GetResponse struct {
	ID       uint     `json:"id"`
	PosterID uint     `json:"poster_id"`
	Content  string   `json:"content"`
	Channels []string `json:"channels"`
}

func (r *GetSerializer) Response() GetResponse {
	myMessageBox := r.c.MustGet("postModel").(Post)
	myChannels := r.c.MustGet("channels").([]Channel)

	var channelContents []string
	for _, channel := range myChannels {
		channelContents = append(channelContents, channel.Content)
	}

	ret := GetResponse{
		ID:       myMessageBox.ID,
		PosterID: myMessageBox.PosterID,
		Content:  myMessageBox.Content,
		Channels: channelContents,
	}
	return ret
}
