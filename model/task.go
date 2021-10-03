package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Title 	 string  `gorm:"type:varchar(100);not null" json:"title"` 	//标题
	Content  string	`gorm:"type:varchar(200)" json:"content"`	//内容
	StartTime int64		//开始时间
	EndTime	  int64		//结束时间
}
