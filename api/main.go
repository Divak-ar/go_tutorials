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

// Model/Schema for course - (goes to model files)

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string  `json:"website"`
}

// fake DB -> (using slice/array to do this -> to store the schema/model type data)

var courses []Course

// middlewares/helpers - goes into separate files
func IsEmpty(c *Course) bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// COntrollers (handling the route logic) - goes in another file

func serveHome(w http.ResponseWriter, r *http.Request){
	// this is boiler plate for writing (w http.ResponseWriter, r *http.Request)
	w.Write([]byte("<h1>Welcome to Golang Tutorials for API Building</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all the courses")

	// setting headers 
	w.Header().Set("Content-Type", "application/json")

	// to read the data from db
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	// we have to take the id from the query params passed in request
	
	w.Header().Set("Content-Type", "application/json")

	// getting the id from request
	params := mux.Vars(r)

	fmt.Printf("The type of req parameters or params is %T and it's value : %v\n\n", params, params)

	// find the matching id -> easier to do with db.findOne(id) query
	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// similar to res.json({"id":"No course exist for this particular id"})
	json.NewEncoder(w).Encode("No Course found with that given ID")
}

// Inserting data into the db(slices here)

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// Passing empty data like {} empty json

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if IsEmpty(&course){
		json.NewEncoder(w).Encode("Empty Json : NO data")
		return
	}

	// generate unique id (convert it into string)
	// append the new course into courses slice

	rand.Seed(time.Now().UnixNano())
	// convrting to string
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// updating the db
	courses = append(courses, course)
	// Returning the response
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse ( w http.ResponseWriter, r *http.Request){
	
}

func main() {
	fmt.Println("Working with APIs in GO 👟👟👟 : (CRUD Operations on building an API)")


}