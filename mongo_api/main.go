package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Divak-ar/mongo_go_api/router"
)

// go get <pkg name>
// additional tools : ctrl + shift + p , type - go tools , then install the tools that you may need for development

func main() {
	fmt.Println("Working with Apis and mongoDB, To get Started : ")
	fmt.Println("https://github.com/mongodb/mongo-go-driver")

	r := router.Router()

	fmt.Println("Launching Server ....")
	log.Fatal(http.ListenAndServe(":4000", r))
}