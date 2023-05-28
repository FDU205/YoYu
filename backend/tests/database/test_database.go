package main

import (
	"YOYU/backend/database"
	"YOYU/backend/users"
)

func main() {
	db := database.TestInit()
	db.AutoMigrate(&users.User{})
	testUser1 := users.User{Username: "zzx", Password: "xxx"}
	testUser2 := users.User{Username: "zzxx", Password: "yyy"}
	err := users.CreateUser(&testUser1)
	if err != nil {
		println(err.Error())
	}

	err = users.CreateUser(&testUser2)
	if err != nil {
		println(err.Error())
	}

}
