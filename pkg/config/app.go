package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	dsn := "root:Averagevs@tcp(127.0.0.1:3306)/cvwowebforum?charset=utf8mb4&parseTime=True&loc=Local"
	if d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		db = d
	}
}

func GetDB() *gorm.DB {
	return db
}
