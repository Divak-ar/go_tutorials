package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Basic File System module : 🔥🔥🔥")

	content := "This is the content of a file..."

	file, err := os.Create("./mygofile.txt")

	if err!=nil{
		panic(err)
		// just shut the program execution
	}

	len, err := io.WriteString(file, content)

	if err!=nil{
		panic(err)
	}

	fmt.Println("io.WriteString() Writes into file with the string and gives its length : ", len)
	// use defer as we can still use the file , by defer the file will close at last after execution of all statements
	defer file.Close()

	readFile(`C:\Users\DIVAKAR\Desktop\go_tutorials\fs\mygofile.txt`)
}

func readFile(filename string){
	// for file maipulation -> ioutile , for creation os 
	databyte, err := ioutil.ReadFile(filename)
	// whenever we read data from online or file it comes in byte format
	if err!=nil{
		panic(err)
	}

	fmt.Println("Text data inside the file is : ", databyte)
	fmt.Println("Text data inside the file is : ", string(databyte))
}