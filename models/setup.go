package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/online_store_db"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Users{}, &Product{}, Order{})

	DB = db
}
