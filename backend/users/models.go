package users

import (
	"YOYU/backend/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"column:username; not null; unique"`
	Password string `gorm:"column:password; not null"`
}

// 设置密码
func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

// 校验密码
func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// 创建新用户
func CreateUser(data interface{}) error {
	db := database.GetDB()
	err := db.Save(data).Error
	return err
}

// 根据条件查询某一用户
func GetUser(condition interface{}) (User, error) {
	db := database.GetDB()
	var userModel User
	err := db.Where(condition).First(&userModel).Error
	return userModel, err
}
