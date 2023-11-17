package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/go-sql-driver/sql"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

// our initial migration function
func initialMigration() {
	//db, err := gorm.Open("sqlite3", "test.db") // db is of type *gorm.DB
	db, err := gorm.Open("mysql", "root:Root@@tcp(127.0.0.1:3306)/curdapi")

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func AllUsers(resp http.ResponseWriter, req *http.Request) {
	// db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("mysql", "root:Root@@tcp(127.0.0.1:3306)/curdapi")
	
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(resp).Encode(users)
}
func newUsers(resp http.ResponseWriter, req *http.Request) {
	// db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("mysql", "root:Root@@tcp(127.0.0.1:3306)/curdapi")

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(req)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	//fmt.Fprintf(resp, "New User Successfully Created")
}

func deleteUser(resp http.ResponseWriter, req *http.Request) {
	// db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("mysql", "root:Root@@tcp(127.0.0.1:3306)/curdapi")

    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(req)
    name := vars["name"]

    var user User
    db.Where("name = ?", name).Find(&user)
    db.Delete(&user)

    fmt.Fprintf(resp, "Successfully Deleted User")
}

func updateUser(resp http.ResponseWriter, req *http.Request) {
	// db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("mysql", "root:Root@@tcp(127.0.0.1:3306)/curdapi")

    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(req)
    name := vars["name"]
    email := vars["email"]

    var user User
    db.Where("name = ?", name).Find(&user)

    user.Email = email

    db.Save(&user)
    fmt.Fprintf(resp, "Successfully Updated User")
}
