package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// to use the github pkg like mux directly :- go mod tidy
// go mod verify :- to verify the integrity of mod files
// go list , go list all , go list -m all :- Gives the dependenices on which our program is dependent
// go list -m -versions github.com/gorilla/mux
// go mod graph :- Gives the graph of all dependencies
// go mod vendor :- vendor folder (related to deployment -- licensne and other things)

func main() {
	fmt.Println("Hello we will deep dive with mod files in GO")
	greeter()

	// In terminal : go mod init github.com/Divak-ar/gomodules 
	// go.mod keeps track of all the dependencies
	// go.sum keeps track of the current version and changes made to them if any (it also verify the legitmicy of the file )

	// gorilla mux -  go get -u github.com/gorilla/mux   for web routing 
	// In new version of >= 1.20 gorilla-mux is no longer required for dynamic routing instead net/http provides it right off the back
	
	// creating a router : const router = express.router() in nodejs
	r := mux.NewRouter()
	// handling a Get request on / route -> router.get("/", middleware , controller ) or router.get("/", async (req, res) => { ...... get somthing from the body await the response , run the ml model or perform some operations(search query) , rreturn the result  })
	r.HandleFunc("/", serveHome).Methods("GET")
	
	// app.listen( PORT , () => { console.log(`Listening to port : ` PORT)}) ; const app = express() and const PORT = 3000 or dotenv()  , const PORT = process.env.PORT || 3000

	log.Fatal(http.ListenAndServe(":4000", r))
	// http://localhost:4000/

}

func greeter(){
	fmt.Println("Hello there mods users")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Golang</h1>"))
}