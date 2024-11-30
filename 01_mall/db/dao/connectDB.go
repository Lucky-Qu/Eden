package dao

import (
	"Eden/01_mall/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB() (db *gorm.DB) {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: config.DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "01_mall",
		},
	})
	if err != nil {
		panic(err)
	}
	return db
}
