package dao

import (
	"to-do-list/repository/db/model"
)

// 执行数据迁移
func migration() {
	// 自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		return
	}
	// DB.Model(&Task{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE")
}
