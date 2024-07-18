package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main(){
	fmt.Println("Handling Get Request 💪💪💪")

	// handling the get request the normal way 
	// PerfromGetRequest()
	// handling the get request using strings pkg ; string pkg provide additional features
	// GetRequestUsingStringsPkg()

	fmt.Println("Handling Post Request 📬📬📬")
	// post handles mainly two types of data : json or urlencoded
	
	// json data
	PerfromPostJsonRequest()

	// urleconded data (Encode means coding the data into particular charset like utf-8/32/64 so that any device world over can read it, as the data sent from the user may not be compatible on the other machine)
	PerfromPostFormRequest()

}

// Captialized first letter for any func means that they can exported and use in other files , perform() becomes private while Perform() becomes public

func PerfromGetRequest(){
	const url = "http://localhost:3000/get"

	res, err := http.Get(url)

	// response is in https format or type
	fmt.Printf("\n Res type: %T", res)
	fmt.Printf("\n Res.Body type: %T", res.Body)
	// response.Body is any end of file signal 


	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("\n", res)
		fmt.Println("Status COde: " , res.StatusCode)
		fmt.Println("Content-Length: ", res.ContentLength)
		fmt.Println("HEader: ", res.Header)
		fmt.Println()
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	// data is in bytes format

	fmt.Println(string(data))
}



func GetRequestUsingStringsPkg(){
	const url = "http://localhost:3000/get"

	res, err := http.Get(url)

	if err != nil{
		fmt.Println("Got an error: ", err)
	}
	fmt.Println()

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	// Building our res string
	var resString strings.Builder

	// using string builder to write the byte data , basically storing it in resString (weird Ik)
	byteCount, _ := resString.Write(data)

	fmt.Println("ByteCount is : ", byteCount)
	fmt.Println("Response is : ", resString.String())
}

func PerfromPostJsonRequest(){
	const url = "http://localhost:3000/post"

	// fake json data/payload / analogy for fetching the user data from req.body 
	// NewReader alllows to create any type of data inside of it, just pass it in that format

	requestBody := strings.NewReader(`
	{
		"name": "GoLang",
		"version": 22.5
	}
	`)

	// normally in post request we get data from user , update the database and return a success msg : here we are creating the post method to send the data and also store it and echo/display it to us

	// post takes (url, contentType, body) hover over Post to see
	res, err := http.Post(url, "application/json", requestBody)

	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	fmt.Println(string(data))
}

func PerfromPostFormRequest(){
	const myurl = "http://localhost:3000/postform"
	
	// creatin a fake form data
	data := url.Values{}

	// injecting values into the form data
	data.Add("firstname: ", "Divakar")
	data.Add("email: ", "divakar@go.dev")

	
	res, err := http.PostForm(myurl, data)

	if err != nil{
		panic(err)
	}
	
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

}