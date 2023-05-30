package messagebox

import (
	"YOYU/backend/database"
	"YOYU/backend/users"
	"bytes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"net/http"
	"net/http/httptest"
	_ "regexp"
)

var MessageBoxRequestTests = []struct {
	init           func(*http.Request)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}{

	//---------------------   Testing for create   ---------------------
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
		"/api/messageBox",
		"POST",
		`{"title": "test1"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":^[1-9]\d*$,"owner_id":^[1-9]\d*$,"title":"test1"},"err_msg":null}`,
		"创建成功",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox",
		"POST",
		`{"title": "test2"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":^[1-9]\d*$,"owner_id":^[1-9]\d*$,"title":"test2"},"err_msg":null}`,
		"创建成功",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox",
		"POST",
		`{"title": "test3"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":^[1-9]\d*$,"owner_id":^[1-9]\d*$,"title":"test3"},"err_msg":null}`,
		"创建成功",
	},
	{
		func(req *http.Request) {},
		"/api/messagebox/create",
		"POST",
		`{}`,
		http.StatusUnprocessableEntity,
		`{"code":1,"err_msg":"参数错误"}`,
		"创建失败，参数错误",
	},
	//---------------------   Testing for get   ---------------------
	{
		func(req *http.Request) {},
		"/api/messagebox/1/1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[{"id":3,"poster_id":1,"content":"test3","visibility":1}\]},"err_msg":null}`,
		"获取帖子3",
	},
	{
		func(req *http.Request) {},
		"/api/messagebox/2/1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[{"id":2,"poster_id":1,"content":"test2","visibility":1}\]},"err_msg":null}`,
		"获取帖子2",
	},
	{
		func(req *http.Request) {},
		"/api/messagebox/1/2",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[{"id":3,"poster_id":1,"content":"test3","visibility":1},{"id":2,"poster_id":1,"content":"test2","visibility":1}\]},"err_msg":null}`,
		"获取帖子3-2",
	},
	{
		func(req *http.Request) {},
		"/api/messagebox/2/2",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[{"id":1,"poster_id":1,"content":"test1","visibility":1}\]},"err_msg":null}`,
		"获取帖子1",
	},
	{
		func(req *http.Request) {},
		"/api/messagebox/3/3",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[\]},"err_msg":null}`,
		"获取帖子3",
	},
	{
		func(req *http.Request) {},
		"/api/messagebox/1/200",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":1,"data":null,"err_msg":"参数错误"}`,
		"获取失败，参数错误",
	},
}

func ResetDB(db *gorm.DB) {
	db.Exec("drop table if exists message_boxes")
	db.Exec("drop table if exists users")
	db.Commit()
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&MessageBox{})
}

func TestMessageBox(t *testing.T) {
	asserts := assert.New(t)

	// 初始化数据库
	test_db := database.TestInit()
	ResetDB(test_db)
	test_db = database.TestInit()
	AutoMigrate(test_db)

	r := gin.New()
	v1 := r.Group("/api")
	// 用户模块
	userG := v1.Group("/user")
	userG.Use(users.AuthMiddleware(false))
	users.UsersRegister(userG)

	// 表白墙模块
	messageBoxG := v1.Group("/messagebox")
	messageBoxG.Use(users.AuthMiddleware(true))
	MessageBoxRegister(messageBoxG)

	var token string
	for i, testData := range MessageBoxRequestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		if i != 0 {
			req.Header.Set("Authorization", "BEARER "+token)
		}
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
		if i == 0 {
			token = w.Body.String()[45:165]
		}
	}
}
