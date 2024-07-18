package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){

	// bufio is used for reading inputs from system(os) stdin (keyboard or any devices)
	reader := bufio.NewReader(os.Stdin)
	// fmt.Scan() 

	// ,_ syntax : since go has no try-catch , here we treat errors like variables that can be stored and returned in case of failure of any func using ,_ syntax with any func (comma error syntax)

	// _, err := func() : the func doesn't return anything of meaning , we are interested in it's successful execution , hence looking for err
	// x, _ := func() : not interested in error only return value
	// x, err := func()
	fmt.Println("Enter a value between 1 to 10: ")
	input, _ := reader.ReadString('\n')
	// '\n' : in readString means read all the user input till the user hits newline (or enter on keyboard)
	fmt.Println("Thanks : ", input)
	fmt.Printf("Type of input : %T", input)

	// COnverting our input var from string to number (conversion) using strconv library

	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil{
		fmt.Println(err)
		// panic(err): in real world this is used to end the execution of the program incase of an error
	}else{
		fmt.Printf("Type of input : %T", numRating)
		// you will get an error(or bug in this case) : stringstrconv.ParseFloat: parsing "7\r\n": invalid syntax it says there is \n going into our input var ... which is true sicnce we have set it for readString() to end on \n, we have to trim this 
	}	

	numRates, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("Type of input : %T and value is %v", numRates, numRates)
		fmt.Println()
	}

	// Stupidity of Go Time module
	presentTime := time.Now()
	fmt.Println(presentTime)

	// Now if you want date only pass into Format("01-02-2006") this specific value 
	fmt.Println(presentTime.Format("01-02-2006"))
	// to get the daay pass monday into format
	fmt.Println(presentTime.Format("Monday"))
	// to get time pass 15:04:05
	fmt.Println(presentTime.Format("15:04:05"))

	// goenv in terminal : gives you all the files for development , GOOS=linux will give you a linux exe file that can run on linux , GOOS windows will give you windows exe file , go run {filename} , go build -> to build exe for main.go (it by default look for main.go) 
	// to get the list of Os and arch(arhictecture for compiled program): go tool dist list

	// Reminder: Go is pass by value and not pass by reference (use pointers for it)

	// Defer statement : defer before a func or statement will make it excute at the end of that func body (so it skips it execution when encountered and execute it at last) .... multiple deffered are executed in LIFO fashion (a stack), reverse order in which they are defered

	fmt.Println()
	defer fmt.Println("Hello - Will executed last as LIFO")
	fmt.Println("World")

	defer fmt.Println("One")
	defer fmt.Println("Two - Will executed before other defer as LIFo")
	fmt.Println("What's the use of defer? ")
	fmt.Println("defer -> allows a function to postpone the execution of a statement until the surrounding function has completed")

	myDefer(6)

	// in defer stack we have 5---1(in myDefer func) , two , one , Hello :- order of output for defer statement
}

func myDefer (num int) {
	for i:= 1; i < num; i++{
		defer fmt.Println(i)
	}
}