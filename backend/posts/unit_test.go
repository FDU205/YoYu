package posts

import (
	"YOYU/backend/database"
	"YOYU/backend/messagebox"
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

var RegisterRequestTests = []struct {
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
		"/api/user/register",
		"POST",
		`{"username": "zzx2", "password": "123456"}`,
		http.StatusOK,
		`{"code":0,"data":{"username":"zzx2","id":2,"token":"[a-zA-Z0-9-_.]{120}"},"err_msg":null}`,
		"注册成功",
	},
}

var MessageRequestTests = []struct {
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
		"/api/messageBox",
		"POST",
		`{"title": "test1"}`,
		http.StatusOK,
		`{"code":0,"data":{"id":1,"owner_id":1,"title":"test1"},"err_msg":null}`,
		"创建提问箱",
	},
}

var PostRequestTests = []struct {
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
		"/api/post",
		"POST",
		`{"message_box_id": 1,"content":"1", "visibility": 1}`,
		http.StatusOK,
		`{"code":0,"data":{"id":1,"poster_id":2,"message_box_id":1,"content":"1","visibility":1},"err_msg":null}`,
		"提问",
	},
	{
		func(req *http.Request) {},
		"/api/post",
		"POST",
		`{"message_box_id": 1,"content":"2", "visibility": 2}`,
		http.StatusOK,
		`{"code":0,"data":{"id":2,"poster_id":2,"message_box_id":1,"content":"2","visibility":2},"err_msg":null}`,
		"提问",
	},
	{
		func(req *http.Request) {},
		"/api/posts?message_box_id=1&page_num=1&page_size=2",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[{"id":2,"poster_id":2,"message_box_id":1,"content":"2","visibility":2},{"id":1,"poster_id":2,"message_box_id":1,"content":"1","visibility":1}\]},"err_msg":null}`,
		"查帖子1-2",
	},
	{
		func(req *http.Request) {},
		"/api/posts?message_box_id=1&page_num=2&page_size=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[{"id":1,"poster_id":2,"message_box_id":1,"content":"1","visibility":1}\]},"err_msg":null}`,
		"查帖子1",
	},
	{
		func(req *http.Request) {},
		"/api/posts?message_box_id=1&page_num=2&page_size=2",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[\]},"err_msg":null}`,
		"查不到帖子",
	},
	{
		func(req *http.Request) {},
		"/api/post/1",
		"DELETE",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":null,"err_msg":null}`,
		"删除帖子1",
	},
	{
		func(req *http.Request) {},
		"/api/posts?message_box_id=1&page_num=1&page_size=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[{"id":2,"poster_id":2,"message_box_id":1,"content":"2","visibility":2}\]},"err_msg":null}`,
		"查帖子2",
	},
	{
		func(req *http.Request) {},
		"/api/mypost?page_num=1&page_size=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[{"id":2,"poster_id":2,"message_box_id":1,"content":"2","visibility":2}\]},"err_msg":null}`,
		"查帖子2",
	},
}

var ChannelRequestTests = []struct {
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
		"/api/post/channel",
		"POST",
		`{"post_id": 2, "content": "追问", "type": 1}`,
		http.StatusOK,
		`{"code":1,"data":null,"err_msg":"无权限追问"}`,
		"无权限追问",
	},
	{
		func(req *http.Request) {},
		"/api/post/channel",
		"POST",
		`{"post_id": 2, "content": "回答", "type": 2}`,
		http.StatusOK,
		`{"code":1,"data":null,"err_msg":"无权限回复"}`,
		"无权限回答",
	},
	{
		func(req *http.Request) {},
		"/api/post/channel",
		"POST",
		`{"post_id": 2, "content": "回答", "type": 2}`,
		http.StatusOK,
		`{"code":0,"data":{"id":1,"post_id":2,"content":"回答","type":2},"err_msg":null}`,
		"回答",
	},
	{
		func(req *http.Request) {},
		"/api/post/channel",
		"POST",
		`{"post_id": 2, "content": "追问", "type": 1}`,
		http.StatusOK,
		`{"code":0,"data":{"id":2,"post_id":2,"content":"追问","type":1},"err_msg":null}`,
		"追问",
	},
	{
		func(req *http.Request) {},
		"/api/post/2",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"id":2,"post_id":2,"post_name":"匿名用户","content":"2","visibility":2,"message_box_id":1,"threads":\[{"id":1,"post_id":2,"content":"回答","type":2},{"id":2,"post_id":2,"content":"追问","type":1}\],"channels":\[\]},"err_msg":null}`,
		"查帖子2",
	},
	{
		func(req *http.Request) {},
		"/api/mypost?page_num=1&page_size=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[{"id":2,"poster_id":2,"message_box_id":1,"content":"2","visibility":2}\]},"err_msg":null}`,
		"查帖子2",
	},
	{
		func(req *http.Request) {},
		"/api/mypost?page_num=1&page_size=1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":\[\]},"err_msg":null}`,
		"查帖子2",
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
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&messagebox.MessageBox{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Channel{})
}

func TestPosts(t *testing.T) {
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
	messagebox.MessageBoxRegister(v1)

	// 帖子模块
	v1.Use(users.AuthMiddleware(true))
	PostRegister(v1)

	// 注册
	var token []string
	for _, testData := range RegisterRequestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
		token = append(token, w.Body.String()[52:172])
	}

	// 创建提问箱
	for _, testData := range MessageRequestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "BEARER "+token[0])
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
	}

	// 提问, 查看问题，删除问题
	for _, testData := range PostRequestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "BEARER "+token[1])
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
	}

	// 回答
	for i, testData := range ChannelRequestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "BEARER "+token[i%2])
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+testData.msg)
	}
}
