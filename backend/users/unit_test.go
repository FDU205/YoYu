package users

import (
	"YOYU/backend/database"
	"YOYU/backend/middlewares"
	"bytes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"net/http"
	"net/http/httptest"
	_ "regexp"
)

func newUser() User {
	return User{
		ID:       2,
		Username: "zzx",
		Password: "",
	}
}

func TestUserModel(t *testing.T) {
	asserts := assert.New(t)

	//Testing UserModel's password feature
	userModel := newUser()
	err := userModel.CheckPassword("")
	asserts.Error(err, "empty password should return err")

	userModel = newUser()
	err = userModel.SetPassword("")
	asserts.Error(err, "empty password can not be set null")

	userModel = newUser()
	err = userModel.SetPassword("123456")
	asserts.NoError(err, "password should be set successful")
	asserts.Len(userModel.Password, 60, "password hash length should be 60")

	err = userModel.CheckPassword("12345")
	asserts.Error(err, "password should be checked and not validated")

	err = userModel.CheckPassword("123456")
	asserts.NoError(err, "password should be checked and validated")
}

var UserRequestTests = []struct {
	init           func(*http.Request)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}{

	//---------------------   Testing for user register   ---------------------
	{
		func(req *http.Request) {},
		"/api/user/register",
		"POST",
		`{"username": "zzx1", "password": "123456"}`,
		http.StatusOK,
		`{"code":0,"data":{"username":"zzx1","id":1,"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"注册成功",
	},
	{
		func(req *http.Request) {},
		"/api/user/register",
		"POST",
		`{"username": "zzx2", "password": "123456"}`,
		http.StatusOK,
		`{"code":0,"data":{"username":"zzx2","id":2,"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"注册成功",
	},
	{
		func(req *http.Request) {},
		"/api/user/register",
		"POST",
		`{"username": "zzx3", "password": "123456"}`,
		http.StatusOK,
		`{"code":0,"data":{"username":"zzx3","id":3,"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"注册成功",
	},
	{
		func(req *http.Request) {},
		"/api/user/register",
		"POST",
		`{"username": "zzx4", "password": "123456"}`,
		http.StatusOK,
		`{"code":0,"data":{"username":"zzx4","id":4,"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"注册成功",
	},
	{
		func(req *http.Request) {},
		"/api/user/register",
		"POST",
		`{"username": "zzx5", "password": "123456"}`,
		http.StatusOK,
		`{"code":0,"data":{"username":"zzx5","id":5,"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"注册成功",
	},
	{
		func(req *http.Request) {},
		"/api/user/register",
		"POST",
		`{"username": "zzx1", "password": "123456"}`,
		http.StatusOK,
		`{"code":1,"data":null,"err_msg":"该用户名已被使用"}`,
		"重复注册",
	},
	{
		func(req *http.Request) {},
		"/api/user/register",
		"POST",
		`{"username": "栈指针栈指针在战争中最最最最最最最最最最最最最最的啊低洼低洼为的", "password": "123456"}`,
		http.StatusOK,
		`{"code":0,"data":{"username":"栈指针栈指针在战争中最最最最最最最最最最最最最最的啊低洼低洼为的","id":7,"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"32位的用户名",
	},
	{
		func(req *http.Request) {},
		"/api/user/register",
		"POST",
		`{"username": "栈指针x栈指针在战争中最最最最最最最最最最最最最最的啊低洼低洼为的", "password": "123456"}`,
		http.StatusUnprocessableEntity,
		`{"code":1,"data":null,"err_msg":"参数错误"}`,
		"33位的用户名",
	},

	//---------------------   Testing for user login   ---------------------
	{
		func(req *http.Request) {},
		"/api/user/login",
		"POST",
		`{"username": "zzx1", "password": "123456"}`,
		http.StatusOK,
		`{"code":0,"data":{"username":"zzx1","id":1,"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"登陆成功",
	},
	{
		func(req *http.Request) {},
		"/api/user/login",
		"POST",
		`{"username": "zzx6", "password": "123456"}`,
		http.StatusOK,
		`{"code":1,"data":null,"err_msg":"用户不存在"}`,
		"用户不存在",
	},
	{
		func(req *http.Request) {},
		"/api/user/login",
		"POST",
		`{"username": "zzx1", "password": "12345"}`,
		http.StatusOK,
		`{"code":1,"data":null,"err_msg":"用户名或密码错误"}`,
		"密码错误",
	},
}

var FollowRequestTests = []struct {
	init           func(*http.Request)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}{
	{
		func(req *http.Request) {},
		"/api/user/follow",
		"POST",
		`{"follow_id":1}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"关注",
	},
	{
		func(req *http.Request) {},
		"/api/user/follow",
		"POST",
		`{"follow_id":2}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"关注",
	},
	{
		func(req *http.Request) {},
		"/api/user/follow",
		"POST",
		`{"follow_id":3}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"关注",
	},
	{
		func(req *http.Request) {},
		"/api/user/follow",
		"POST",
		`{"follow_id":4}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"关注",
	},
	{
		func(req *http.Request) {},
		"/api/user/follow",
		"POST",
		`{"follow_id":5}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"关注",
	},
}

var unFollowRequestTests = []struct {
	init           func(*http.Request)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}{
	{
		func(req *http.Request) {},
		"/api/user/unfollow",
		"DELETE",
		`{"follow_id":1}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"取消关注",
	},
	{
		func(req *http.Request) {},
		"/api/user/unfollow",
		"DELETE",
		`{"follow_id":2}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"取消关注",
	},
	{
		func(req *http.Request) {},
		"/api/user/unfollow",
		"DELETE",
		`{"follow_id":3}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"取消关注",
	},
	{
		func(req *http.Request) {},
		"/api/user/unfollow",
		"DELETE",
		`{"follow_id":4}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"取消关注",
	},
	{
		func(req *http.Request) {},
		"/api/user/unfollow",
		"DELETE",
		`{"follow_id":5}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"取消关注",
	},
}

var FollowCountRequestTests = []struct {
	init           func(*http.Request)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}{
	{
		func(req *http.Request) {},
		"/api/user/followcount",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"count":5,"err_msg":null}`,
		"关注数量",
	},
	{
		func(req *http.Request) {},
		"/api/user/isfollow?follow_id=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"err_msg":null,"yes":true}`,
		"关注数量",
	},
	{
		func(req *http.Request) {},
		"/api/user/fanscount",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"count":5,"err_msg":null}`,
		"关注数量",
	},
	{
		func(req *http.Request) {},
		"/api/user/followcount",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"count":0,"err_msg":null}`,
		"关注数量",
	},
	{
		func(req *http.Request) {},
		"/api/user/fanscount",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"count":0,"err_msg":null}`,
		"关注数量",
	},
	{
		func(req *http.Request) {},
		"/api/user/isfollow?follow_id=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"err_msg":null,"yes":false}`,
		"关注数量",
	},
}

var FollowListRequestTests = []struct {
	init           func(*http.Request)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}{
	{
		func(req *http.Request) {},
		"/api/user/followlist?page_num=1&page_size=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"count":5,"err_msg":null}`,
		"关注数量",
	},
	{
		func(req *http.Request) {},
		"/api/user/fanslist?page_num=2&page_size=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"count":5,"err_msg":null}`,
		"关注数量",
	},
	{
		func(req *http.Request) {},
		"/api/user/fanslist?page_num=2&page_size=6",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"count":5,"err_msg":null}`,
		"关注数量",
	},
}

func ResetDB(db *gorm.DB) {
	db.Exec("drop table if exists channels")
	db.Exec("drop table if exists posts")
	db.Exec("drop table if exists message_boxes")
	db.Exec("drop table if exists walls")
	db.Exec("drop table if exists followers")
	db.Exec("drop table if exists users")
	db.Commit()
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Follower{})
}

func TestUsers(t *testing.T) {
	asserts := assert.New(t)

	// 初始化数据库
	test_db := database.TestInit()
	ResetDB(test_db)
	test_db = database.TestInit()
	AutoMigrate(test_db)

	// 注册路由
	r := gin.New()
	v1 := r.Group("/api")
	userG := v1.Group("/user")
	userG.Use(middlewares.AuthMiddleware(false))
	UsersRegister(userG)
	userG.Use(middlewares.AuthMiddleware(true))
	FollowsRegister(userG)

	// 登陆，注册
	var token []string
	for i, testData := range UserRequestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
		if i < 5 {
			token = append(token, w.Body.String()[52:172])
		}
	}

	// 关注
	for id := 0; id < 5; id = id + 1 {
		for _, testData := range FollowRequestTests {
			bodyData := testData.bodyData
			req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "BEARER "+token[id])
			asserts.NoError(err)

			testData.init(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
			asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
		}
	}

	// 查看数量
	for id := 0; id < 5; id = id + 1 {
		for _, testData := range FollowCountRequestTests[0:3] {
			bodyData := testData.bodyData
			req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "BEARER "+token[id])
			asserts.NoError(err)

			testData.init(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
			asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
		}
	}

	// 关注列表
	for id := 0; id < 5; id = id + 1 {
		for _, testData := range FollowListRequestTests {
			bodyData := testData.bodyData
			req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "BEARER "+token[id])
			asserts.NoError(err)

			testData.init(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
			//asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+w.Body.String())
		}
	}

	// 取关
	for id := 0; id < 5; id = id + 1 {
		for _, testData := range unFollowRequestTests {
			bodyData := testData.bodyData
			req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "BEARER "+token[id])
			asserts.NoError(err)

			testData.init(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
			asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
		}
	}

	// 查看数量
	for id := 0; id < 5; id = id + 1 {
		for _, testData := range FollowCountRequestTests[3:] {
			bodyData := testData.bodyData
			req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "BEARER "+token[id])
			asserts.NoError(err)

			testData.init(req)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
			asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
		}
	}
}
