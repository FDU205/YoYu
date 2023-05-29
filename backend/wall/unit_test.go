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

	//---------------------   Testing for create   ---------------------
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
		"/api/wall/create",
		"POST",
		`{"content":"test", visibility: 0}`,
		http.StatusOK,
		`{"code":0,"err_msg":null}`,
		"创建成功",
	},
	{
		func(req *http.Request) {},
		"/api/wall/create",
		"POST",
		`{"content":"test", visibility: 2}`,
		http.StatusUnprocessableEntity,
		`{"code":0,"err_msg":"参数错误"}`,
		"创建失败，参数错误",
	},
	{
		func(req *http.Request) {},
		"/api/wall/create",
		"POST",
		`{}`,
		http.StatusUnprocessableEntity,
		`{"code":0,"err_msg":"参数错误"}`,
		"创建失败，参数错误",
	},
	//---------------------   Testing for get   ---------------------
	{
		func(req *http.Request) {},
		"/api/wall/1/1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":{"content":"test","id":1,"poster_id":1,"visibility":0}},"err_msg":null}`,
		"获取成功",
	},
	{
		func(req *http.Request) {},
		"/api/wall/2/1",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":{"posts":null},"err_msg":null}`,
		"获取成功",
	},
	{
		func(req *http.Request) {},
		"/api/wall/1/",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":null,"err_msg":"参数错误"}`,
		"获取失败，参数错误",
	},
	{
		func(req *http.Request) {},
		"/api/wall/1/200",
		"GET",
		`{}`,
		http.StatusOK,
		`{"code":0,"data":null,"err_msg":"参数错误"}`,
		"获取失败，参数错误",
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

	var token string
	for i, testData := range WallRequestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		if i != 0 {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		asserts.NoError(err)

		testData.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Regexp(testData.responseRegexg, w.Body.String(), "Response Content - "+w.Body.String())
		if i == 0 {
			token = w.Body.String()[27:164]
		}
	}
}
