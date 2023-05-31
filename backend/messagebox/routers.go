package messagebox

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 将MessageBox模块的功能注册进框架
func MessageBoxRegister(router *gin.RouterGroup) {
	router.POST("/messageBox", Create)
	router.GET("/messageBoxes", Search)
	router.GET("/messageBox/:id", Get)
	router.PUT("/messageBox/:id", Update)
	router.DELETE("/messageBox/:id", Delete)
}

// 创建提问箱
func Create(c *gin.Context) {
	messageBoxValidator := NewMessageBoxValidator()
	if err := messageBoxValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	if err := MessageBoxCreate(&messageBoxValidator.MessageBoxModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": messageBoxValidator.MessageBoxModel})
}

// 根据提问箱ID获取提问箱
func Get(c *gin.Context) {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	messageBox, err := MessageBoxGetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "找不到提问箱", "data": nil})
		return
	}

	c.Set("messageBoxModel", messageBox)
	//TODO: 获取posts
	posts := []uint{}
	c.Set("posts", posts)

	serializer := GetSerializer{c}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()})
}

// 查询提问箱
func Search(c *gin.Context) {
	title := c.Query("title")
	ownerID_str := c.Query("owner")
	page_num_str := c.Query("page_num")
	page_size_str := c.Query("page_size")

	var ownerID int
	if ownerID_str != "" {
		var err error
		ownerID, err = strconv.Atoi(ownerID_str)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
			return
		}
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

	messageBoxes, err := MessageBoxSearch(title, uint(ownerID), page_num, page_size)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.Set("messageBoxes", messageBoxes)
	serializer := SearchSerializer{c}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()})
}

// 根据提问箱ID删除提问箱
func Delete(c *gin.Context) {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	if err := MessageBoxDeleteByID(uint(id), c.MustGet("userID").(uint)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 2, "err_msg": "删除失败", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": nil})
}

// 根据提问箱ID更新提问箱
func Update(c *gin.Context) {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": "参数错误", "data": nil})
		return
	}

	messageBoxValidator := NewMessageBoxValidator()
	if err := messageBoxValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": "参数错误"})
		return
	}

	messageBoxValidator.MessageBoxModel.ID = uint(id)
	if err := MessageBoxUpdateByID(messageBoxValidator.MessageBoxModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 2, "err_msg": "更新失败", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": messageBoxValidator.MessageBoxModel})
}
