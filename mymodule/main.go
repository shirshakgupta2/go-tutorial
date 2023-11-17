package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")

	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	http.ListenAndServe(":4545", r)

	// log.Fatal(http.ListenAndServe(":9090",r))

}

func greeter() {
	fmt.Println(" Hey the mod users")
}

func serverHome(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<h1> Welcome to golang series</h1>\n"))
}
