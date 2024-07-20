package main

import (
	"encoding/json"
	"fmt"
	"log"
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


	var course Course

	// getting the data from user and storing it in course
	_ = json.NewDecoder(r.Body).Decode(&course)

	// Passing empty data like {} empty json
	if IsEmpty(&course){
		json.NewEncoder(w).Encode("Empty Json : NO data")
		return
	}

	// generate unique id (convert it into string)
	// append the new course into courses slice

	fmt.Println(course.CourseName)
	fmt.Printf("The data sent is: %v" , course)

	// handling duplicate courses (don't add it to the db , check and return)
	for _ , c := range courses{
		if c.CourseName == course.CourseName{
			json.NewEncoder(w).Encode("This course is already present, try updating it's value instead of creating another")
			return;
		}
	}

	rand.Seed(time.Now().UnixNano())
	// convrting to string
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// updating the db
	courses = append(courses, course)

	// Returning the response . res.send("msg") or res.status(200).send("msg") in nodejs
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse ( w http.ResponseWriter, r *http.Request){
	// update the course based on unique id - findOneandUpadte(in sql and mongoose) , since we are using any db here
	// we will iterate over slices , get the specific id , make changes , then update the slice with new values
	// how to remake the slice after update -> courses = append(course[:index], course[index+1:]...) , it updates the courses db and removes the eleement with specific index/id


	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	// use mux.Vars(r) to get all the values in request params  (here the courseID)
	params := mux.Vars(r)

	// loop, find the element with specific id and remove the value , get the new value from request body and add it to the db

	for index, course := range courses {
		// courses is our fake db (slice to store the data)
		if course.CourseId == params["id"]{
			// removing that id from courses
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			// getting the updated values from body
			_ = json.NewDecoder(r.Body).Decode(&course)
			// giving the id and the result of values is taken from the body
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			// id found and db updated just return from here
			return
		}
	}

	// Send a response when id is not found 
	json.NewEncoder(w).Encode("COurse of this id doesn't exist")
}

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses{
		if course.CourseId == params["id"] {
			// removing that element
			courses = append(courses[:index], courses[index+1:]...)
			break;
		}
	}

	json.NewEncoder(w).Encode("the course is deleted")

	// return backend response in html format, for that we have to define the string in some variable , here AddForm is variable 
	// if we get some form data by the user then good else we display this 
	// for this : set  w.Header().Set("Content-Type", "text/html; charset=utf-8")

	url := r.FormValue("url")
    if url == "" {
        fmt.Fprint(w, AddForm)
        return
    }
}

func main() {
	fmt.Println("Working with APIs in GO 👟👟👟 : (CRUD Operations on building an API)")

	r := mux.NewRouter()

	// seeding/filling the db with some test data

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{Fullname: "Divakar", Website: "dp.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN", CoursePrice: 599, Author: &Author{Fullname: "Divakar", Website: "mern.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	// whatever we pass here as params query -> {id} the same name has to used when defining the controller , so params["id"] will be used in controller where params := mux.Vars(r) , r is request -> *http.Request type data
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	// sending/creating new data -> Post (must be encrypted and protected)
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	// for updating use put 
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listening to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// routing in golang with MUX

// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

// Restrict the request handler to specific hostnames or subdomains.
// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

// Restrict the request handler to http/https.
// r.HandleFunc("/secure", SecureHandler).Schemes("https")
// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

// Add Queries 
// r.HandleFunc("/authors", authorBook).Queries("surname", "{surname}")
