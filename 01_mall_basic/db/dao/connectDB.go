package dao

import (
	"Eden/01_mall_basic/config"
	"Eden/01_mall_basic/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// init 链接数据库并将db存储到包中,并确定数据库中的表是否存在，如果不存在则创建
func init() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: config.DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "01_mall_",
		},
	})
	if err != nil {
		panic(err)
	}
	if !db.Migrator().HasTable(&model.Product{}) {
		err = db.Migrator().CreateTable(&model.Product{})
		if err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.User{}) {
		err = db.Migrator().CreateTable(&model.User{})
		if err != nil {
			panic(err)
		}
	}
}
