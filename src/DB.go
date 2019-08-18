package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	SingletonDB	*gorm.DB
	onceDB		sync.Once
)

func getDB() *gorm.DB {
	onceDB.Do(func() {
		SingletonDB = connectDB()
	})

	return SingletonDB
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root@/sirChat?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
