package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// execution starts from main as always , the compiler look for main func , rest func are defined in main or out side and then imported here for use

func main(){
	fmt.Println("Learning Go for DevOps and backend!");

	// In go any variable defined has to be used : rule introduce for design simplicity , else will give errors/warnings 

	var num int;
	// by default the int will be casted on int64 or int32 depending on the system architecture
	var overflow uint16 = 32767 + 1;

	var decimal float32;

	var str string;

	// type casting (explicit)
	var res float32 = float32(num) * decimal

	fmt.Println()
	fmt.Println("The number is : ", num, "this is runtime error , int overflow(max limit for int16 is 32767 add 1 we get -ve limit) : ", overflow);
	fmt.Println("the decimal number: ", decimal, "cool dude all in float32 : ", res);
	fmt.Println("integer division are rounded down ", 3/2 , "result get converted into higher data type ", 3/2.0);
	fmt.Println("By default all variables are intialized to 0 (including rune), if string then an empty string - '' : ", str);

	// string : multiple lines use `` , single line "xyz\nyas", we can concatenate string

	var str1 string = "blah blah";
	var str2 string = " yeah hooo";
	
	fmt.Println()
	fmt.Println(str1+str2, "length of str2 (hover cursor over len to read it's output varies as per data types) : ", len(str1))
	
	fmt.Println()
	fmt.Println("for gamma symbol (alt 226 on right numberpad): ", len("Γ"), " for normal letters: ", len("Z"));
	fmt.Println()
	fmt.Println("since go uses utf-8 character set, letters outside of ascii will take more bytes and len gives no. of bytes ..... it's same incase of vanilla ascii characters : to count length while using some fancy symbol use a unicode/utf8 libraray,")

	fmt.Println("\n the length of string: ", utf8.RuneCountInString("Γ"))


	// rune are characters (char) in go
	var rune2 rune = 'a'
	fmt.Println(rune2)

	var torf bool = false;
	fmt.Println(torf, " : default value for boolean is false")

	// fancy way to define variable (data type is inferred by the compiler) btw go is statically typed like java,c++ (the variables defined with data types will store the same data) unlike in Javascript where let can store arrays, number, string, etc . Dynamically typed example : Python

	var good = 34;
	fmt.Println(good)

	// more fancy way : just omit var and use := in place of =
	better := "cool"
	fmt.Println(better)

	a1, a2 := 4 , 5
	fmt.Println("Defining multiple variable in same line : ", a1 * a2)

	// const is for value that will remain the same throughout the program , they have to be initialized explicitly
	const pi = 22/7.0
	fmt.Println(pi)

	// Dealing with functions in GO : functions are defined by func keyword
	printMe()

	n1 , n2 := 4 , 1
	// since we are returning some value , we have to store it , it is better we declare our var with data type as it makes it more readable as compare := synatx

	var div int = divideMe(n1, n2);
	if(div == n2){
		fmt.Println("Division not possible");
	}else{
		fmt.Println("Result: ", div)
	}

	// var k int ;

	// taking input : 
	// fmt.Scan(&k) :- Stops when user gives a space/newline as input 
	// fmt.Scanln(&k) ;- Stops at new line , 
	// fmt.Scanf(%v , &k) ;- Stops when stops receiving the data types input declared : %v is for value

	// fmt.Println("Enter the value of k : ")
	// fmt.Scanln(&k);
	// fmt.Println("The value you entered: ", k)

	// func can give multiple returns

	fmt.Println(n1 , n2)
	var b1, b2 , b3 = calculator(n1, n2)
	fmt.Println("The Sum is: ", b1, "\nThe Subtraction is: ", b2, "\nOn Multiply: ", b3 )
	fmt.Printf("The sum is : %v , The multiply is : %v", b1, b3)
	fmt.Println();

	// A design pattern in go is to encounter error : go throws error if encountered . It can be captured in error type , default value is ni
	// nil is similar to null but In Go, nil is a predeclared identifier representing the zero value for various types, including pointers, slices, maps, channels, interfaces, and functionsl
	var err error ;
	fmt.Println(err);

	var a3,a4 int
	// import errors package to use error
	fmt.Scanln(&a3, &a4);
	
	if a4 == 0 {
		err = errors.New("Error in division: Cannot divide by Zero for god's sake")
		fmt.Println(err)
	}else{
		fmt.Println("Divide: " , a3/a4);
	}

	// More menaingful use (in a func)

	var result , errs = intDivision(a3, a4)

	// now if no error occurs in the code then err should not be nil else it will have "Cannot be divided by zero error : Logical error" value
	if errs != nil {
		fmt.Println("This one is from the func :- ")
		fmt.Println(errs.Error())
		
	}else{
		fmt.Println("Whatever division: ", result)
	}

}

func printMe(){
	fmt.Println("Hello from printMe() Function")
}

// return type is defined before {body} and after (parameter) , here only one return of int type

func divideMe(n1 int, n2 int) int {
	if n2 != 0 {
		return (n1/n2);
	}else{
		return n2
	}
}

func calculator(a1 int, a2 int) (int, int , int){
	return a1+a2, a1-a2, a1*a2;
}

func intDivision(x int, y int) (int, error){
	var err error
	if y == 0{
		err = errors.New("cannot be divided by zero error : Logical error")
		return 0, err
	}

	return x/y , err
}
