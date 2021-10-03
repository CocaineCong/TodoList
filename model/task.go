package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	"to-do-list/cache"
)

//任务模型
type Task struct {
	gorm.Model
	UserId 		  uint
	Title         string
	Status        int  `gorm:"default:'0'"`
	Content       string `gorm:"size:1000"`
	StartTime 	  time.Time
	EndTime 	  time.Time
}

func (Task *Task) View() uint64 {
	//增加点击数
 	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(Task.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//AddView
func (Task *Task) AddView() {
	cache.RedisClient.Incr(cache.ProductViewKey(Task.ID))	//增加视频点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.ID)))	//增加排行点击数

}

