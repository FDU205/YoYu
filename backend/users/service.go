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
