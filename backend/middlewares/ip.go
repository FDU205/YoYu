package middlewares

import (
	"YOYU/backend/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisDb *redis.Client

// 连接到redis
func InitRedis() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// 防火墙中间件
func RateMiddleware(c *gin.Context) {
	// 预设时间内刷新key为IP(c.ClientIP())为0
	err := redisDb.SetNX(c.ClientIP(), 0, common.DUR_TIME).Err()

	// 每次访问，这个IP的对应的值加一
	redisDb.Incr(c.ClientIP())
	if err != nil {
		panic(err)
	}

	// 获取IP访问的次数
	var val int
	val, err = redisDb.Get(c.ClientIP()).Int()
	if err != nil {
		panic(err)
	}
	// 如果大于预设次数，返回403
	if val > common.REQUEST_COUNT {
		c.Abort()
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"data": "请求的太快了， 喝杯咖啡休息一下吧",
		})
		return
	} else {
		// 到下一个中间件
		c.Next()
	}
}
