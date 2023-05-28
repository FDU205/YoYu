package main

import (
	"YOYU/backend/database"
	"YOYU/backend/users"

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

	r.Run("127.0.0.1:8080")
}
