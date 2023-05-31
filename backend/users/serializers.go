package users

import (
	"YOYU/backend/utils"

	"github.com/gin-gonic/gin"
)

type RegisterSerializer struct {
	c *gin.Context
}

type RegisterResponse struct {
	Username string `json:"username"`
	UserId   uint   `json:"id"`
	Token    string `json:"token"`
}

func (r *RegisterSerializer) Response() RegisterResponse {
	myUserModel := r.c.MustGet("userModel").(User)
	user := RegisterResponse{
		Username: myUserModel.Username,
		UserId:   myUserModel.ID,
		Token:    utils.GenToken(myUserModel.ID),
	}
	return user
}

type LoginSerializer struct {
	c *gin.Context
}

type LoginResponse struct {
	Username string `json:"username"`
	UserId   uint   `json:"id"`
	Token    string `json:"token"`
}

func (LS *LoginSerializer) Response() LoginResponse {
	userModel := LS.c.MustGet("userModel").(User)
	user := LoginResponse{
		Username: userModel.Username,
		UserId:   userModel.ID,
		Token:    utils.GenToken(userModel.ID),
	}
	return user
}

type FollowSerializer struct {
	c *gin.Context
}

type Followin struct {
	ID   uint   `json:"user_id"`
	Name string `json:"username"`
}

type FollowResponse struct {
	Follows []Followin `json:"follows"`
}

func (FS *FollowSerializer) Response() FollowResponse {
	myFollow := FS.c.MustGet("follows").([]User)
	var myFollowin []Followin
	for _, follow := range myFollow {
		myFollowin = append(myFollowin, Followin{
			ID:   follow.ID,
			Name: follow.Username,
		})
	}

	ret := FollowResponse{
		Follows: myFollowin,
	}

	return ret
}

type FansSerializer struct {
	c *gin.Context
}

type Fansin struct {
	ID   uint   `json:"user_id"`
	Name string `json:"username"`
}

type FansResponse struct {
	Fans []Fansin `json:"fans"`
}

func (FS *FansSerializer) Response() FansResponse {
	myFans := FS.c.MustGet("fans").([]User)
	var myFansin []Fansin
	for _, fan := range myFans {
		myFansin = append(myFansin, Fansin{
			ID:   fan.ID,
			Name: fan.Username,
		})
	}

	ret := FansResponse{
		Fans: myFansin,
	}

	return ret
}
