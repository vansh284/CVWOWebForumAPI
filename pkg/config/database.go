package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	envMap = GetEnvMap()
	dsn := envMap["DSN"]
	if d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		db = d
	}
}

func GetDB() *gorm.DB {
	return db
}
