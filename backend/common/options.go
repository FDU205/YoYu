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