package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for courses -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware,helper-file
func (c *Course) IsEmpty() bool { // method is a part of structure which is course
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

//controllers -file

// SERVE HOME ROUTE
func serveHome(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<h1>WELCOME TO SERVE HOME</h1>"))
}

func getAllCourses(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Get all courses")
	resp.Header().Set("Content-Type", "application/json") //set header content type
	json.NewEncoder(resp).Encode(courses)                 // NewEncoder() we will pass what is the writer    
	             //what ever we send in encode() will be given bake as the response to the json

}

func getOneCourse(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Get one course")
	resp.Header().Set("Content-Type", "application/json") //set header content type

	// grab id from request
	params := mux.Vars(req) //all the variable that is coming in
	fmt.Printf("the params: %v and the type of params is %T\n", params, params)

	// loop through the courses and find matching id and then return the response
	for _, course := range courses { // courses will have key value pairs where values here we use id
		if course.CourseId == params["id"] {
			json.NewEncoder(resp).Encode(course)
			json.NewEncoder(resp).Encode(course.Author)
			return
		}
	}
	json.NewEncoder(resp).Encode(courses)
	json.NewEncoder(resp).Encode("no courses found  with given ID")
	//return
}

func createOneCourse(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Create one course")
	resp.Header().Set("Content-Type", "application/json") //set header content type

	//if body is empty
	if req.Body == nil {
		json.NewEncoder(resp).Encode("Please send a data to create a new course")

	}
	//create the course object
	var course Course

	_ = json.NewDecoder(req.Body).Decode(&course)

	//what about the data that is{}
	if course.IsEmpty() {
		json.NewEncoder(resp).Encode("Please send some data into json that you send to create a new course")
		json.NewEncoder(resp).Encode(course)
		json.NewEncoder(resp).Encode("the courses are empty")

		return
	}
	//TODO:check if title is a duplicate or not
	//loop through coureses,if titke mathces with any of the course.coursenamethen encode say its the same course that you are giving


	//Generate a  unique ID, Convert id to String
	//Append course into courses
	rand.Seed(time.Now().UnixNano()) //this will give the unique no

	// to convert the integer value to a string
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(resp).Encode(course)
	return

}

func updateOneCourse(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Create one course that wants to be updated")
	resp.Header().Set("Content-Type", "application/json")

	//first-garb id from req
	params := mux.Vars(req)

	//loop through courses and then remove and then add back to course where we have the id
	for index, course := range courses{
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(req.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(resp).Encode(course)
			json.NewEncoder(resp).Encode("the update is completed form the above course id")

			return

		}

	}
	//TODO:send a response if id not found,or if body request is empty,or id=f json is{}

}
func deleteOneCourse(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Create one course that wants to be updated")
	resp.Header().Set("Content-Type", "application/json")

	//first-garb id from req
	params := mux.Vars(req)

	//loop through courses and then remove where we have the id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(resp).Encode(" Course id Deleted!")
			break
		}
	}

	//TODO: if id not found, or if body request is empty
}

func main() {
	fmt.Println("API-we are making CRUD with fake data base in slice")
	r := mux.NewRouter()

	//seeding data base
	courses = append(courses, Course{"1", "ReactJS", 234, &Author{"Shirshak", "gupta"}})
	courses = append(courses, Course{"2", "JAVA", 545, &Author{"yuvraj", "patil"}})
	courses = append(courses, Course{"3", "NodeJS", 256, &Author{"Akshay", "devidas"}})

	//routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")

	//as we are expecting the id in params we give id as a payloads {2}
	r.HandleFunc("/getcourse/{id}", getOneCourse).Methods("GET") //here the url will be localhost:9090/getcourse/2

	r.HandleFunc("/createcourse", createOneCourse).Methods("POST")
	r.HandleFunc("/updatecourse/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/deletecourse/{id}", deleteOneCourse).Methods("DELETE")

	//creata and Listen the port
	http.ListenAndServe(":9090", r)

}
