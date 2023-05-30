package main

import (
	"YOYU/backend/database"
	"YOYU/backend/users"
	"YOYU/backend/wall"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
}

func main() {

	db := database.Init()
	Migrate(db)

	r := gin.Default()

	v1 := r.Group("/api")

	// 用户模块
	userG := v1.Group("/user")
	users.UsersRegister(userG)

	// 表白墙模块
	wallG := v1.Group("/wall")
	wallG.Use(users.AuthMiddleware(true))
	wall.WallRegister(wallG)

	r.Run("127.0.0.1:8080")
}
