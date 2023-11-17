package router

import (
	"net/http"
	"github.com/shirshak/search/search_stocks"
	"github.com/gorilla/mux"
)

func InitRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/searchstock", searchstock.SearchStock)
	http.ListenAndServe(":9898", r)
}
