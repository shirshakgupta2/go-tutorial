package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// we have to define one global variable, which have info about the database.and for that we have to import "gorm.io/gorm"
var Database *gorm.DB

// we have to create urlDSN variable inside in the following format
// var urlDSN = "username:password@tcp(localhost:3306)/NameOfDatabase" where 3306 is the portno of mysql
// var urlDSN = "host=localhost user=postgres password=Root dbname=curdsql port=9920 sslmode=disable TimeZone=Asia/Shanghai" // "username:password@tcp(localhost:3306)/NameOfDatabase" where 3306 is the portno of mysql
var urlDSN = "root:Root@tcp(localhost:3306)/curdsql" // "username:password@tcp(localhost:3306)/NameOfDatabase" where 3306 is the portno of mysql

var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("connection failed !!")
	}
	Database.AutoMigrate(&Movies{})

}
