package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/book_store_go?charset=utf8mb4&parseTime=True&loc=Local") // d = gorm.Open....
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() * gorm.DB{
	return db
}

