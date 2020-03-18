package common

import (
	"github.com/jinzhu/gorm"
)


func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root@/go_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		 panic("Error connecting to database")
	}
	return db
}

