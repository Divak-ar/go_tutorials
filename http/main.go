package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const myurl = "https://www.udemy.com/join/signup-popup/?locale=en_US&response_type=html&return_link=https%3A%2F%2Fwww.udemy.com%2Fcourse%2Fbecome-a-certified-web-developer%2F%3FcouponCode%3DTHANKSLEARNER24&source=course_landing_page&ref=&next=%2Fgift%2Fbecome-a-certified-web-developer%2F%3FcouponCode%3DTHANKSLEARNER24"

func main(){
	fmt.Println("Reading webpage content : 📖📖📖")

	response, err := http.Get(myurl)

	checkError(err)

	fmt.Printf("Response is of Type : %T", response)

	// In go , it is our responsibility to close the connection, so always when you open it then use defer there to close it in case if forget to close it , defer will take care of it  (so it closes at the end after all the func have been performed), why should we close the connection : some services like aws charges based on session (connection time) usage while keeping it open while not using it utilizes resources 
	defer response.Body.Close()

	// Reading the response
	dataBytes, err := ioutil.ReadAll(response.Body)

	checkError(err)

	content := string(dataBytes)

	// content: throws out the whole html page content with code --- Great for reverse engineering a webpage
	fmt.Println(content)



	fmt.Println()
	fmt.Println()
	fmt.Println("------------ Handling URL in GO ---------")

	// Parsing the url
	result, err := url.Parse(myurl)

	checkError(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of Query params is : %T\n", qparams)
	// type are url.Values :- Basically a key:value pairs (a dictionary)

	fmt.Println("returns parameters from the query (?) portion of url : " , qparams["return_link"])
	fmt.Printf("\n\n")
	fmt.Println("Params with there key : values")

	for val, val2 := range qparams{
		fmt.Println(val, " : ", val2)
	}

	fmt.Println()
	fmt.Println()


	// To create your own url 

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "www.geeksforgeeks.org",
		Path: "/building-basic-chrome-extension",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

	
}

func checkError(err error){
	if err != nil{
		fmt.Println("Can't connect to the website , error 404")
	}
}