package users

import (
	"YOYU/backend/database"
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
		`{"code":0,"data":{"username":"zzx1","token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
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

	//---------------------   Testing for user login   ---------------------
	{
		func(req *http.Request) {},
		"/api/user/login",
		"POST",
		`{"username": "zzx1", "password": "123456"}`,
		http.StatusOK,
		`{"code":0,"data":{"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"登陆成功",
	},
	{
		func(req *http.Request) {},
		"/api/user/login",
		"POST",
		`{"username": "zzx2", "password": "123456"}`,
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

func ResetDB(db *gorm.DB) {
	db.Exec("drop table if exists walls")
	db.Exec("drop table if exists users")
	db.Commit()
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func TestUsers(t *testing.T) {
	asserts := assert.New(t)

	// 初始化数据库
	test_db := database.TestInit()
	ResetDB(test_db)
	test_db = database.TestInit()
	AutoMigrate(test_db)

	r := gin.New()
	v1 := r.Group("/api")
	UsersRegister(v1.Group("/user"))
	for _, testData := range UserRequestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+w.Body.String())
	}
}
