package main

import (
	"YOYU/backend/database"
	"YOYU/backend/messagebox"
	"YOYU/backend/posts"
	"YOYU/backend/users"
	"YOYU/backend/wall"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&users.Follower{})
	db.AutoMigrate(&wall.Wall{})
	db.AutoMigrate(&messagebox.MessageBox{})
	db.AutoMigrate(&posts.Post{})
	db.AutoMigrate(&posts.Channel{})
}

func main() {

	// 初始化数据库
	db := database.Init()
	Migrate(db)

	r := gin.Default()

	v1 := r.Group("/api")

	// 用户模块
	userG := v1.Group("/user")
	userG.Use(users.AuthMiddleware(false))
	users.UsersRegister(userG)
	userG.Use(users.AuthMiddleware(true))
	users.FollowsRegister(userG)

	// 表白墙模块
	wallG := v1.Group("/wall")
	wallG.Use(users.AuthMiddleware(true))
	wall.WallRegister(wallG)

	// 提问箱模块
	v1.Use(users.AuthMiddleware(true))
	messagebox.MessageBoxRegister(v1)

	// 帖子模块
	v1.Use(users.AuthMiddleware(true))
	posts.PostRegister(v1)

	r.Run("192.168.1.109:8080")
}
