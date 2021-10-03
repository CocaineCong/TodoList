package cache

import (
	"to-do-list/pkg/logging"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

//RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

//Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64) //TODO 这里记得了！！
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}

