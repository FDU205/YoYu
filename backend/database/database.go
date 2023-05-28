package database

import (
	"YOYU/backend/common"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// 打开数据库并保存引用至全局变量
func Init() *gorm.DB {
	db, err := gorm.Open(mysql.Open(common.DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	sqldb, _ := db.DB()
	// 设置最大空连接数
	sqldb.SetMaxIdleConns(10)
	DB = db
	return DB
}

// 该函数用于获得已连接的数据库
func GetDB() *gorm.DB {
	return DB
}

func TestInit() *gorm.DB {
	dsn := "yoyu:123456@tcp(127.0.0.1:3306)/test_yoyu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	sqldb, _ := db.DB()
	// 设置最大空连接数
	sqldb.SetMaxIdleConns(10)
	DB = db
	return DB
}
