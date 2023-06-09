/*该文件定义了整个项目所使用的常量*/
package common

import "time"

// token的参数
const TOKEN_KEY = "@a secret key to generate YOYU token!$" // 加盐
const EXP = time.Hour * 24                                 // token24小时后过期

// database的参数
// 可自行修改用户名
const DSN = "yoyu:123456@tcp(127.0.0.1:3306)/yoyu?charset=utf8mb4&parseTime=True&loc=Local"

// users的参数
const RANDOM_SECRET = "123456"

// middlewares的参数
const DUR_TIME = 1 * time.Second // ip拦截的时间
const REQUEST_COUNT = 50         // 一定时间内的请求次数上限

// cache的参数
const CACHE_EXP = time.Minute * 2  // cache过期时间
const CACHE_PURG = time.Minute * 5 // cache清理时间
