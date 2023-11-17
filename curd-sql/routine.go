package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// we write a new fun to RountingHandler()

func RountingHandler() {
	r := mux.NewRouter()

	r.HandleFunc("/", test).Methods("POST")
	
	r.HandleFunc("/createMovie", CreateMovie).Methods("POST")

	r.HandleFunc("/Movies", GetMovie).Methods("GET")

	r.HandleFunc("/getMovie/{id}", GetMovieById).Methods("GET")

	r.HandleFunc("/updateMovie/{id}", UpdateMovie).Methods("PUT")

	r.HandleFunc("/deletemovie/{id}", DeleteMovie).Methods("DELETE")

	log.Println("hello world")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
	
}
