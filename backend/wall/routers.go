package wall

import (
	"YOYU/backend/common"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var WallCache *cache.Cache

// 将Wall模块的功能注册进框架
func WallRegister(router *gin.RouterGroup) {
	WallCache = cache.New(common.CACHE_EXP, common.CACHE_PURG)
	router.POST("/create", Create)
	router.GET("", Get)
	router.GET("/mywall", GetMyWall)
}

// 创建表白信息
func Create(c *gin.Context) {
	wallValidator := NewWallValidator()
	if err := wallValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1, "err_msg": "参数错误"})
		return
	}

	if err := WallCreate(&wallValidator.WallModel); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "err_msg": err.Error()})
		// 清空cache
		WallCache.Flush()
		return
	}

	// 清空cache
	WallCache.Flush()
	c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": nil})
}

// 返回今日所有表白信息
func Get(c *gin.Context) {
	// cache中找
	if ret, found := WallCache.Get(c.Request.URL.String()); found {
		c.JSON(http.StatusOK, ret.(gin.H))
		return
	}
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

	wall, err := GetWallByPage(time.Now().Format("2006-01-02"), page_num, page_size)

	var ret gin.H
	if err != nil {
		ret = gin.H{"code": 1, "err_msg": err.Error(), "data": nil}
	} else {
		c.Set("wallModel", wall)
		serializer := GetSerializer{c}
		ret = gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()}
	}

	// 存cache
	WallCache.Set(c.Request.URL.String(), ret, cache.DefaultExpiration)
	c.JSON(http.StatusOK, ret)
}

// 返回我的表白信息
func GetMyWall(c *gin.Context) {
	// cache中找
	if ret, found := WallCache.Get(c.Request.URL.String() + c.GetHeader("Authorization")); found {
		c.JSON(http.StatusOK, ret.(gin.H))
		return
	}

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

	UserID := c.MustGet("userID").(uint)
	wall, err := MyWallGet(UserID, page_num, page_size)

	var ret gin.H
	if err != nil {
		ret = gin.H{"code": 1, "err_msg": err.Error(), "data": nil}
	} else {
		c.Set("wallModel", wall)
		serializer := GetSerializer{c}
		ret = gin.H{"code": 0, "err_msg": nil, "data": serializer.Response()}
	}

	// 存cache
	WallCache.Set(c.Request.URL.String()+c.GetHeader("Authorization"), ret, cache.DefaultExpiration)
	c.JSON(http.StatusOK, ret)
}
