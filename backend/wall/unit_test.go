package wall

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

var WallRequestTests = []struct {
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
		`{"code":0,"data":{"username":"zzx1","token":"[a-zA-Z0-9-_.]{137}"},"err_msg":null}`,
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
		`{"code":0,"data":{"token":"[a-zA-Z0-9-_.]{137}"},"err_msg":null}`,
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
	db.Exec("drop table if exists users")
	db.Commit()
}

func TestUsers(t *testing.T) {
	asserts := assert.New(t)
	test_db := database.TestInit()
	ResetDB(test_db)
	test_db = database.TestInit()
	test_db.AutoMigrate(&Wall{})

	r := gin.New()
	v1 := r.Group("/api")
	WallRegister(v1.Group("/wall"))
	for _, testData := range WallRequestTests {
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
