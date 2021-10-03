package model

import (
	"cmall/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

//商品模型
type Product struct {
	gorm.Model
	CategoryID    int
	CategoryName  string
	Title         string
	Status        int
	Info          string `gorm:"size:1000"`
	BossID        int
	BossName      string
}

func (Product *Product) View() uint64 {
	//增加视频点击数
 	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(Product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//AddView
func (Product *Product) AddView() {
	//增加视频点击数
	cache.RedisClient.Incr(cache.ProductViewKey(Product.ID))
	//增加排行点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Product.ID)))
}

//AddElecRank 增加加点排行点击数
func (Product *Product) AddTaskRank() {
	//增加家电排汗点击数
	cache.RedisClient.ZIncrBy(cache.TaskRank, 1, strconv.Itoa(int(Product.ID)))
}
//
////AddAcceRank 增加配件排行点击数
//func (Product *Product) AddAcceRank() {
//	//增加配件排行点击数
//	cache.RedisClient.ZIncrBy(cache.AccessoryRank, 1, strconv.Itoa(int(Product.ID)))
//}
