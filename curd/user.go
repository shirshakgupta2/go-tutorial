package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func initialMigration() {
	//db, err := gorm.Open("sqlite3", "test.db") // db is of type *gorm.DB
	// db, err := gorm.Open("")

	dsn := "host=localhost user= password=Root dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}
