package model

import "github.com/jinzhu/gorm"

type Status struct {
	gorm.Model
	NotBegin string   //未开始
	NowIng string	  //正在做
	Achieve string    //完成
}