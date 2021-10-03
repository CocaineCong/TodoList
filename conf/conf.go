package conf

import (
	"cmall/cache"
	"cmall/model"
	"cmall/pkg/logging"
	"github.com/joho/godotenv"
	"strings"
)

const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "remember"
)

func Init() {
	_ = godotenv.Load() //从本地读取环境变量
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logging.Info(err) //日志内容
		panic(err)
	}
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
	cache.Redis()
}
