package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 将Users模块的功能注册进框架
func UsersRegister(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
}

// 注册关注模块的功能
// 这些是要鉴权之后才能使用的
func FollowsRegister(router *gin.RouterGroup) {
	router.POST("/follow", Follow)
	router.DELETE("/unfollow", UnFollow)
	router.GET("/isfollow", IsFollow)
	router.GET("/followlist", GetFollowList)
	router.GET("/followcount", GetFollowCount)
	router.GET("/fanslist", GetFansList)
	router.GET("/fanscount", GetFansCount)
}

// 注册
func Register(c *gin.Context) {
	userValidator := NewRegisterValidator()
	if err := userValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	if err := UserRegister(&userValidator.UserModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.Set("userModel", userValidator.UserModel)
	serializer := RegisterSerializer{c}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()})
}

// 登陆
func Login(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}
	userModel, err := UserLogin(loginValidator.UserModel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.Set("userModel", userModel)
	serializer := LoginSerializer{c}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()})
}

// 关注
func Follow(c *gin.Context) {
	followValidator := NewFollowValidator()
	if err := followValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": err.Error()})
		return
	}

	if err := UserFollow(followValidator.FollowModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "关注失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil})
}

// 取关
func UnFollow(c *gin.Context) {
	followValidator := NewFollowValidator()
	if err := followValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": err.Error()})
		return
	}

	if err := UserUnFollow(followValidator.FollowModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "取关失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil})
}

// 是否关注
func IsFollow(c *gin.Context) {
	followID_str := c.Query("follow_id")

	followID, err := strconv.Atoi(followID_str)
	if err != nil || followID < 1 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "yes": nil})
		return
	}

	userID := c.MustGet("userID").(uint)

	yes := IsFollowing(&Follower{FollowingID: uint(followID), FollowedByID: userID})
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "yes": yes})
}

// 获取关注列表
func GetFollowList(c *gin.Context) {
	page_num_str := c.Query("page_num")
	page_size_str := c.Query("page_size")
	if page_num_str == "" || page_size_str == "" {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	page_num, err := strconv.Atoi(page_num_str)
	if err != nil || page_num < 1 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	page_size, err := strconv.Atoi(page_size_str)
	if err != nil || page_size < 1 || page_size > 100 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	id := c.MustGet("userID").(uint)
	follows := FollowListGet(id, page_num, page_size)

	c.Set("follows", follows)
	serializer := FollowSerializer{c}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()})
}

// 获取关注数
func GetFollowCount(c *gin.Context) {
	id := c.MustGet("userID").(uint)
	count, err := FollowCountGet(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "错误", "count": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "count": count})
}

// 获取粉丝列表
func GetFansList(c *gin.Context) {
	page_num_str := c.Query("page_num")
	page_size_str := c.Query("page_size")
	if page_num_str == "" || page_size_str == "" {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	page_num, err := strconv.Atoi(page_num_str)
	if err != nil || page_num < 1 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	page_size, err := strconv.Atoi(page_size_str)
	if err != nil || page_size < 1 || page_size > 100 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	id := c.MustGet("userID").(uint)
	fans := FansListGet(id, page_num, page_size)

	c.Set("fans", fans)
	serializer := FansSerializer{c}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()})
}

// 获取粉丝数
func GetFansCount(c *gin.Context) {
	id := c.MustGet("userID").(uint)
	count, err := FansCountGet(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "错误", "count": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "count": count})
}
