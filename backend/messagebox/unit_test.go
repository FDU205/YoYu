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
		`{"code":0,"data":{"username":"zzx1","id":1,"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"注册成功",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox",
		"POST",
		`{"title": "test1"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":1,"owner_id":1,"title":"test1"},"err_msg":null}`,
		"创建成功",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox",
		"POST",
		`{"title": "test2"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":2,"owner_id":1,"title":"test2"},"err_msg":null}`,
		"创建成功",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox",
		"POST",
		`{"title": "test3"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":3,"owner_id":1,"title":"test3"},"err_msg":null}`,
		"创建成功",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox",
		"POST",
		`{"title": "test4"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":4,"owner_id":1,"title":"test4"},"err_msg":null}`,
		"创建成功",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox",
		"POST",
		`{"title": "test5"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":5,"owner_id":1,"title":"test5"},"err_msg":null}`,
		"创建成功",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox/3",
		"DELETE",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":null,"err_msg":null}`,
		"删除提问箱",
	},
	//---------------------   Testing for get   ---------------------
	{
		func(req *http.Request) {},
		"/api/messageBox/1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"id":1,"owner_id":1,"title":"test1","posts":\[\]},"err_msg":null}`,
		"获取提问箱1",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox/3",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":1,"data":null,"err_msg":"找不到提问箱"}`,
		"获取提问箱3",
	},
	{
		func(req *http.Request) {},
		"/api/messageBoxes?page_num=1&page_size=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"messageBoxes":\[{"id":5,"owner_id":1,"title":"test5","owner_name":"zzx1"}\]},"err_msg":null}`,
		"获取提问箱5",
	},
	{
		func(req *http.Request) {},
		"/api/messageBoxes?page_num=2&page_size=2",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"messageBoxes":\[{"id":2,"owner_id":1,"title":"test2","owner_name":"zzx1"},{"id":1,"owner_id":1,"title":"test1","owner_name":"zzx1"}\]},"err_msg":null}`,
		"获取提问箱2-1",
	},
	{
		func(req *http.Request) {},
		"/api/messageBoxes?page_num=2&page_size=6",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"messageBoxes":\[\]},"err_msg":null}`,
		"获取不到提问箱",
	},
	{
		func(req *http.Request) {},
		"/api/messageBoxes?page_num=1&page_size=1&title=test",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"messageBoxes":\[{"id":5,"owner_id":1,"title":"test5","owner_name":"zzx1"}\]},"err_msg":null}`,
		"查找test提问箱",
	},
	{
		func(req *http.Request) {},
		"/api/messageBoxes?page_num=1&page_size=1&owner=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"messageBoxes":\[{"id":5,"owner_id":1,"title":"test5","owner_name":"zzx1"}\]},"err_msg":null}`,
		"找到用户1的提问箱",
	},
	{
		func(req *http.Request) {},
		"/api/messageBoxes?page_num=1&page_size=1&owner=2",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"messageBoxes":\[\]},"err_msg":null}`,
		"找不到提问箱",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox/1",
		"PUT",
		`{"title":"test10"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":1,"owner_id":1,"title":"test10"},"err_msg":null}`,
		"修改提问箱标题",
	},
	{
		func(req *http.Request) {},
		"/api/messageBox/1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"id":1,"owner_id":1,"title":"test10","posts":\[\]},"err_msg":null}`,
		"获取提问箱1",
	},
}

func ResetDB(db *gorm.DB) {
	db.Exec("drop table if exists message_boxes")
	db.Exec("drop table if exists walls")
	db.Exec("drop table if exists followers")
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

	// 提问箱模块
	v1.Use(users.AuthMiddleware(true))
	MessageBoxRegister(v1)

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
			token = w.Body.String()[52:172]
		}
	}
}
