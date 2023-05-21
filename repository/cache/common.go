package cache

import (
	"strconv"

	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"

	conf "to-do-list/config"
)

var RedisClient *redis.Client

// 初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(conf.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
