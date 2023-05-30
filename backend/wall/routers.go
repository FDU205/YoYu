package wall

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 将Wall模块的功能注册进框架
func WallRegister(router *gin.RouterGroup) {
	router.POST("/create", Create)
	router.GET("/:page_num/:page_size", Get)
}

// 创建表白信息
func Create(c *gin.Context) {
	wallValidator := NewWallValidator()
	if err := wallValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": "参数错误"})
		return
	}

	if err := WallCreate(&wallValidator.WallModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil})
}

// 返回今日所有表白信息
func Get(c *gin.Context) {
	page_num_str := c.Param("page_num")
	page_size_str := c.Param("page_size")

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

	wall, err := GetWallByPage(time.Now().Format("2006-01-02"), page_num, page_size)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error(), "data": nil})
		return
	}

	serializer := GetSerializer{wall}
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil, "data": serializer})
}
