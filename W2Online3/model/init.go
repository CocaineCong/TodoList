package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	//db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
	//	DisableForeignKeyConstraintWhenMigrating: true,// 外键约束
	//	SkipDefaultTransaction: true, 	// 禁用默认事务（提高运行速度）
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
	//	},
	//})
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)       //默认不加复数s
	db.DB().SetMaxIdleConns(20)  //设置连接池，空闲
	db.DB().SetMaxOpenConns(100) //打开
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//db.Model(&Boss{}).AddForeignKey("ProductID", "Product(ProductID)", "RESTRICT", "RESTRICT")
	DB = db
	migration()
}
