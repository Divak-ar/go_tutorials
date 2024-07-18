package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

)

func main(){
	fmt.Println("Get/Post req in go 🆒🆒🆒")

	// to get the response from rapidAPI , we have to pass the api key in headers  

	godotenv.Load(".env")
	PerformGETRequest()  
	
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	 
	// get the value
	return os.Getenv(key)
  }

func PerformGETRequest(){
	const url string = `https://api.thedogapi.com/v1/images/search?limit=2`

	// pass the key to get the value from env file
	dotenv := goDotEnvVariable("DOG_API")
	
	req, err := http.NewRequest("GET", url, nil)

	// pass the key as the request header for authentication
	req.Header.Add("x-api-key", dotenv)


	checkError(err)

	// send the request 
	res, err := http.DefaultClient.Do(req)

	checkError(err)

	// close the request 
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Printf("\nThe type of data is : %T\n", res)

	fmt.Println("\n\n After Converting to String the data looks like: ", string(body))

	res_body := string(body)
	fmt.Printf("\nThe type of data is : %T\n\n", res_body)

	fmt.Println("Splitting the data based on []\t" , strings.Split(res_body, "[]"))
}

func checkError(err error) {
	if err != nil{
		fmt.Println("Error Occured while get req : ", err)
	}
}