package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dns := "shunux:PnSnNn@11@tcp(127.0.0.1:3306)/crud"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Blog{})
	DB = db
}
