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
	users.UsersRegister(v1.Group("/user"))
	wall.WallRegister(v1.Group("/wall"))

	r.Run("127.0.0.1:8080")
}
