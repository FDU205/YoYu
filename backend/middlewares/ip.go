package middlewares

import (
	"YOYU/backend/common"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisDb *redis.Client
var whiteList []string
var blackList []string

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

	// 读黑白名单
	f1, err := os.Open("./middlewares/white.txt")
	if err != nil {
		return err
	}
	f2, err := os.Open("./middlewares/black.txt")
	if err != nil {
		return err
	}

	defer f1.Close()
	defer f2.Close()

	r := bufio.NewReader(f1)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading white file %s", err)
			break
		}
		whiteList = append(whiteList, line)
	}

	r = bufio.NewReader(f2)
	for {

		line, _, err := r.ReadLine()
		if err != nil {
			fmt.Printf("error reading blackfile %s", err)
			break
		}
		blackList = append(blackList, string(line))
	}

	return nil
}

func InWhiteList(ip string) bool {
	for _, s := range whiteList {
		if s == ip {
			return true
		}
	}
	return false
}

func InBlackList(ip string) bool {
	for _, s := range blackList {
		if s == ip {
			return true
		}
	}
	return false
}

// 防火墙中间件
func RateMiddleware(c *gin.Context) {
	ip := c.ClientIP()
	// 如果在黑白名单内
	if InWhiteList(ip) {
		c.Next()
		return
	} else if InBlackList(ip) {
		c.Abort()
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"err_msg": "403Fobidden!\n请求的太快了, 喝杯咖啡休息一下吧",
			"data":    nil,
		})
		return
	}

	// 预设时间内刷新key为IP(ip)为0
	err := redisDb.SetNX(ip, 0, common.DUR_TIME).Err()

	// 每次访问，这个IP的对应的值加一
	redisDb.Incr(ip)
	if err != nil {
		panic(err)
	}

	// 获取IP访问的次数
	var val int
	val, err = redisDb.Get(ip).Int()
	if err != nil {
		panic(err)
	}
	// 如果大于预设次数，返回403
	if val > common.REQUEST_COUNT {
		c.Abort()
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"err_msg": "403Fobidden!\n请求的太快了, 喝杯咖啡休息一下吧",
			"data":    nil,
		})
		return
	} else {
		// 到下一个中间件
		c.Next()
	}
}
