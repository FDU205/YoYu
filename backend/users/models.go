package users

import (
	"YOYU/backend/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 用户表
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

// 关注表
type Follower struct {
	gorm.Model
	Following    User
	FollowingID  uint
	FollowedBy   User
	FollowedByID uint
}

// u关注v
func Following(data interface{}) error {
	db := database.GetDB()
	err := db.Save(data).Error
	return err
}

// u取消关注v
func UnFollowing(data interface{}) error {
	db := database.GetDB()
	err := db.Where(data).Delete(&Follower{}).Error
	return err
}

// u是否关注v
func IsFollowing(data interface{}) bool {
	db := database.GetDB()
	var follow Follower
	db.Where(data).First(&follow)
	return follow.ID != 0
}

// 获取关注列表
func FollowingsList(id uint, offset int, limit int) []User {
	db := database.GetDB()
	tx := db.Begin()
	var follows []Follower
	var followings []User
	tx.Where(&Follower{
		FollowedByID: id,
	}).Order("updated_at desc").Limit(limit).Offset(offset).Find(&follows)
	for _, follow := range follows {
		var userModel User
		db.Where(&User{ID: follow.FollowingID}).First(&userModel)
		followings = append(followings, userModel)
	}
	tx.Commit()
	return followings
}

// 获取关注数
func FollowingCount(id uint) (int64, error) {
	db := database.GetDB()
	var ret int64
	err := db.Model(&Follower{}).Where(&Follower{
		FollowedByID: id,
	}).Count(&ret).Error
	return ret, err
}

// 获取粉丝列表
func FansList(id uint, offset int, limit int) []User {
	db := database.GetDB()
	tx := db.Begin()
	var fans []Follower
	var fansUser []User
	tx.Where(&Follower{
		FollowingID: id,
	}).Order("updated_at desc").Limit(limit).Offset(offset).Find(&fans)
	for _, fan := range fans {
		var userModel User
		db.Where(&User{ID: fan.FollowedByID}).First(&userModel)
		fansUser = append(fansUser, userModel)
	}
	tx.Commit()
	return fansUser
}

// 获取粉丝数
func FansCount(id uint) (int64, error) {
	db := database.GetDB()
	var ret int64
	err := db.Model(&Follower{}).Where(&Follower{
		FollowingID: id,
	}).Count(&ret).Error
	return ret, err
}
