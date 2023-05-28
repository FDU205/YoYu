package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 调用gin框架的binding, 将前端发来的参数邦定至对应结构体
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}
