package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloworld).Methods("GET")
	r.HandleFunc("/users", AllUsers).Methods("GET")
	r.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	r.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	r.HandleFunc("/user/{name}/{email}", newUsers).Methods("POST")

	http.ListenAndServe(":9090", r)
}

func main() {
	fmt.Println("GO ORM ")
	handleRequests()
	initialMigration()

}
