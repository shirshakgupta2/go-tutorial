package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// creatng Handler Functions
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	//first of all we have to set our imcoming Response data type(w) to json format
	w.Header().Set("Content-Type", "application/json")
	var mov Movies
	json.NewDecoder(r.Body).Decode(&mov)

	//to storing emp to db
	Database.Create(&mov)

	//After creating and storing values into db , we return back the values in we stored to the frontend part in json format

}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var AllMov []Movies
	//to fetch all records form db and storeing it into AllEmp
	Database.Find(&AllMov)
	//to find name => Database.db.Where("name = ?", "yash").First(&user)
	//Database.Where("emp_name = ?", "yas").Find(&AllEmp)

	//return json format to frontend
	json.NewEncoder(w).Encode(&AllMov)

}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mov Movies

	params := mux.Vars(r)

	Database.First(&mov, params["id"])
	json.NewEncoder(w).Encode(&mov)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mov Movies
	params := mux.Vars(r)
	// fmt.Fprintf(w, "the Movie cost that i  will be updating : %v\n", )
	// fmt.Fprintf(w, "the Movie name that i  will be updating : %v\n", movv.MovieName)
	// fmt.Fprintf(w, "the params : %v\n", params)

	// Database.Delete(&movv, params["id"])
	// Database.First(&movv, params["id"]).Where("movie_id = ?", params["id"])
	// Database.Model(&movv).Where("movie_id = ?", params["id"]).Update(movv)
	// Database.First(&mov,params["id"])

	// Database.Model(Movies{}).Where("movie_id = ?", params["id"]).Updates(&mov)
	// fmt.Fprintf(w, "the Movie id that i  will be updating : %v\n", mov.MovieName)
	// json.NewDecoder(r.Body).Decode(&mov)
	// fmt.Fprintf(w, "the Movie id that i  will be updating : %v\n", mov.MovieName)
	// fmt.Fprintf(w, "the parameter id that i looked for: %v\n", params["id"])

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	Database.Model(&mov).Where("movie_id = ?", id).Updates(mov)

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&mov)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update user to update the user

	Database.Save(&mov)
	json.NewEncoder(w).Encode(mov)

}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mov Movies
	params := mux.Vars(r)
	Database.Delete(&mov, params["id"])
	json.NewEncoder(w).Encode("The Data Releted to This Id get Deleted Succesfully !")

}
