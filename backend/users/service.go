package users

import "errors"

func UserRegister(userModel interface{}) error {
	err := CreateUser(userModel)
	if err != nil {
		return errors.New("该用户名已被使用")
	}
	return nil
}

func UserLogin(userModel User) (User, error) {
	findUser, err := GetUser(&User{Username: userModel.Username})
	if err != nil {
		return findUser, errors.New("用户不存在")
	}

	err = findUser.CheckPassword(userModel.Password)
	if err != nil {
		return findUser, errors.New("用户名或密码错误")
	}
	return findUser, nil
}

func UserFollow(followModel Follower) error {
	return Following(&followModel)
}

func UserUnFollow(followModel Follower) error {
	return UnFollowing(&followModel)
}

func FollowListGet(id uint, pageNum int, pageSize int) []User {
	offset := (pageNum - 1) * pageSize
	limit := pageSize
	user := FollowingsList(id, offset, limit)
	return user
}

func FollowCountGet(id uint) (int64, error) {
	return FollowingCount(id)
}

func FansListGet(id uint, pageNum int, pageSize int) []User {
	offset := (pageNum - 1) * pageSize
	limit := pageSize
	user := FansList(id, offset, limit)
	return user
}

func FansCountGet(id uint) (int64, error) {
	return FansCount(id)
}
