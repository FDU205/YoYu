package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 将Post模块的功能注册进框架
func PostRegister(router *gin.RouterGroup) {
	router.POST("/post", Create)
	router.POST("/post/channel", CreateAnswer)
	router.GET("/posts", Search)
	router.GET("/post/:id", Get)
	router.GET("/mypost", GetMyPost)
	router.DELETE("/post/:id", Delete)
}

// 创建帖子
func Create(c *gin.Context) {
	postValidator := NewPostValidator()
	if err := postValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	if err := PostCreate(&postValidator.PostModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": postValidator.PostModel})
}

// 根据帖子ID获取帖子
func Get(c *gin.Context) {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	channels, post, err := PostGetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.Set("postModel", post)
	c.Set("channels", channels)

	serializer := GetSerializer{c}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()})
}

// 查询帖子
func Search(c *gin.Context) {
	message_box_id_str := c.Query("message_box_id")
	page_num_str := c.Query("page_num")
	page_size_str := c.Query("page_size")

	message_box_id, err := strconv.Atoi(message_box_id_str)
	if err != nil {
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

	posts, err := PostSearch(uint(message_box_id), page_num, page_size)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	serializer := SearchSerializer{posts}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer})
}

// 根据帖子ID删除帖子
func Delete(c *gin.Context) {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	if err := PostDeleteByID(uint(id), c.MustGet("userID").(uint)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 2, "err_msg": "删除失败", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": nil})
}

// 创建回复
func CreateAnswer(c *gin.Context) {
	channelValidator := NewChannelValidator()
	if err := channelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": "参数错误"})
		return
	}

	if err := ChannelCreate(channelValidator.OwnerID, &channelValidator.ChannelModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": channelValidator.ChannelModel})
}

// 查询帖子
func GetMyPost(c *gin.Context) {
	page_num_str := c.Query("page_num")
	page_size_str := c.Query("page_size")

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

	userID := c.MustGet("userID").(uint)
	posts, err := MyPostGet(userID, page_num, page_size)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	serializer := SearchSerializer{posts}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer})
}
