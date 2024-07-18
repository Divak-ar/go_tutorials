package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"`
	Price int	`json:"price"`
	Platform string	
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

// `json:"Alias_Name_For_Key"` basically give some other name to key here
// `json:"-"` : ignore this field when converting in marshal{eg: password}"`
// `json:"tags,omitempty"`: if there is nil value , doesn't put it in output (omit it)
// `json: "Key"`: will give not compatible error .... `json:" Key"` : error.... blah blah 
// format `json:"Key"`

func main() {
	fmt.Println("Creating JSON Data using struct as go doesn't deal with OOPS")

	// EncodeJson()


	fmt.Println("Dealing with JSON data from backend/API ╰(*°▽°*)╯ ╰(*°▽°*)╯ ╰(*°▽°*)╯")

	DecodeJson()


}

func EncodeJson(){

	webCourse := []course{
		{"REact", 2999, "xyz.com", "shshs22", []string{"web-dev", "js", "frontend"}},
		{"Vue.js", 3999, "xyz.com", "224dshs22", nil},
		{"MERN", 5999, "xyz.com", "sagdu232hs22", []string{"web-dev", "js", "fullstack"}},
	}

	// package this data as JSON data using Marshall in json enconding

	// finalJson, err := json.Marshal(webCourse)

	// more readable than Marshal, takes prefix and the value on which the data should be indented
	finalJson, err := json.MarshalIndent(webCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%s\n\n", finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	 {
                "coursename": "REact",
                "price": 2999,
                "Platform": "xyz.com",
                "tags": [
                        "web-dev",
                        "js",
                        "frontend"
                ]
        }
	`)

	
	var webCourse course

	// verifying the data format
	checkValid := json.Valid(jsonDataFromWeb)

	if(checkValid){
		fmt.Println("data was of valid json type")
	}else{
		fmt.Println("data was of not of valid json type")
	}

	json.Unmarshal(jsonDataFromWeb, &webCourse)

	fmt.Printf("%#v\n", webCourse)


	// there are some cases where we just want to add data to key value pair

	// In json key is always : string but the value could be anything -> slice(array), int, string . So we use interface now so that we can populate it later based on the data type we receive in value
	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	fmt.Printf("\n%#v\n", myData)

	for k, v := range myData{
		fmt.Printf("The key : %v  / the value : %v / and the Type of value is: %T\n", k, v, v)
	}
}