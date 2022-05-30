package redisdb

import (
	"fmt"
	"go/conf/utils"
	"go/pkg/mylog"

	//"go/pkg/mylog"
	"time"

	//"log"
	//"go/pkg/utils"
	"github.com/go-redis/redis"
)

var log = mylog.NewLog("Warning")

var r *redis.Client

func init() {
	// r = redis.NewClient(&redis.Options{
	// 	Addr:     utils.GetRedishost(), // redis地址
	// 	Password: utils.GetRedispwd(),  // redis密码，没有则留空
	// 	DB:       utils.GetRedisdb(),   // 默认数据库，默认是0
	// })

	// //通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	// _, err := r.Ping().Result()
	// if err != nil {
	// 	log.Error(err.Error())
	// }
	go func() {
		for {
			r = redis.NewClient(&redis.Options{
				Addr:     utils.GetRedishost(), // redis地址
				Password: utils.GetRedispwd(),  // redis密码，没有则留空
				DB:       utils.GetRedisdb(),   // 默认数据库，默认是0
			})

			//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
			_, err := r.Ping().Result()
			if err != nil {
				log.Error(err.Error())
				time.Sleep(15 * time.Second)
				fmt.Println("123")
			} else {
				break
			}
		}

	}()
}
func Setkey(key string, value string, overtime int) bool {
	err := r.Set(key, value, time.Duration(overtime)*time.Second).Err() //存入redis 设置过期
	//	r.SetNX()
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}
func Getvalue(key string) string {
	val, err := r.Get(key).Result()
	// 判断查询是否出错
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return val
}
