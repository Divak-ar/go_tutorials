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
		err = errors.New("error in division: Cannot divide by Zero for god's sake")
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



	// -------------------------------------------------- //

	// In go break after each statement is implicitly declared , so no need to write (functions like if-elseif-else ladder)

	switch  {
	case errs!=nil:
		fmt.Println(err.Error())
	case errs==nil:
		fmt.Println("result : " , result)
	}

	// other way(The general way) to use switch to perform multiple task for different value of a expression

	var expression rune
	fmt.Scan(&expression)

	val := expression - '9'
	fmt.Println("val is: ", val)

	switch errs {
	case nil:
		fmt.Println(errs)
	default:
		fmt.Println("Yeah")
	}

	var value int;
	fmt.Scanln(&value)

	switch value {
	case 1:
	  fmt.Println("Value is 1")
	case 2, 3:  // Match multiple values
	  fmt.Println("Value is 2 or 3")
	fallthrough // Fall through to the next case even if a match occurred above
	case 4:
	  fmt.Println("Value is 4 or less (including 1, 2, 3)")
	default:
	  fmt.Println("Value is greater than 4")
	}

	// In this example, if value is 2, both the case 2, 3 and case 4 blocks will be executed due to fallthrough.






	// ------------------------------- Arrays , Slices , Map and Loops

	const alen int = 5;

	var int32Arr [alen]int32;

	for i := 0; i < alen; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		_,eror := fmt.Scan(&int32Arr[i])

		if eror != nil {
            fmt.Println("Error reading input:", err)
            break // Exit the loop if there's an error
        }
	}

	fmt.Println(int32Arr[0], " The Second Element: ", int32Arr[1])
	fmt.Println("Index Slicing: " , int32Arr[0: alen-2])
	fmt.Println("Getting the Adress: ", &int32Arr[1], " ", &int32Arr[2])


	// intArr := [3]int32{4,7,9} is same as intArr := [...] int32 {6,3,5,7}

	intArr := [...] int32 {6,3,5,7}
	fmt.Println(intArr)
	fmt.Println("Arrays are of fixed size , there size has to be predefined with a constant has it size can't be assigned/changed on runtime")
	fmt.Println("To create an array of dynamic size , whose size we assign at runtime use Slices")

	// Slices : 
	var intSlice []int32 = []int32{2,3,5}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println("Cap is a function that gives capacity of the Slice")

	// append is used to increment length (a new array is made to accomodate 7 just like ArrayList in java or Vectors in c++) the size incremented here Slices in Go don't directly append to the pre-existing array. They create a new underlying array with a doubled capacity (up to a certain limit) when necessary to accommodate additional elements.

	intSlice = append(intSlice, 7)
	// 7 is added at the end of the intSlice 
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println()

	// can append another slice to the pre-existing one (adding mutiple values) using sliceName... operaator
	var intSlice2 []int32 = []int32 {43,23,56}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice2), cap(intSlice2))
	fmt.Println()

	// another way to make slice is using make fn , it takes datatype, length of slice, capacity of slice(optional)
	var intSlice3 []int32 = make([]int32 , 3, 6)
	// adding values
	intSlice3 = append(intSlice3, 10, 20, 30) 
	fmt.Println(intSlice3)
	fmt.Printf("The length is %v with capacity %v", len(intSlice3), cap(intSlice3))
	fmt.Println()



	// Maps just like in c++ and Java (In python a dictionary tho fucntionality is different) : {key: value} pair , here key:string and value:uint16
	var myMap map[string]uint16 =  make(map[string]uint16)
	fmt.Println(myMap)

	var myMap2 = map[string]uint16{"Adam":21, "Sarah":26, "Jackson Pollock": 35}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	// map return a (optional) boolean t/f depending on if the element is present int the map or not
	var age, ok = myMap2["Jordan"]
	if ok{
		fmt.Println("The age is %v: ", age)
	}else{
		fmt.Println("Invalid Name, not present in the map")
	}

	// To delete a value
	delete(myMap2, "Jackson Pollock")
	fmt.Println(myMap2)

	// To iterate : loops (range)

	for name, age := range myMap2{
		// while iterating over ma : there is no order to which data is stored (like unordered map) the values will be given at random and it will change when run again (2 output for this iteration will not be same {in most of the time})
		fmt.Print("Name : ", name)
		fmt.Print(" and Age : ", age)
		fmt.Println()
	}

	// iterating over arrays and slices
	for i, v := range intSlice3{
		fmt.Printf("Index: %v and Value is %v \n", i, v)
	}

	// Go doesn't have while loop, it can be acheived through for loop , eg : This is while loop in GO
	i := 0;
	for i < 4{
		fmt.Println(i)
		// i = i+1;
		i += 1;
	}

	// Or instead of above we can add break keyword (if condition) to make it behave like while loop
	fmt.Println()
	for {
		if i >= 7{
			break
		}
		fmt.Println(i)
		i = i+1;
	}

	fmt.Println()
	// Normal for Loopz
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	}


	// Deep dive into strings and rune

	var str5 string = "resumé"
	fmt.Println(str5)

	var indexed = str5[0];
	fmt.Println(indexed)
	fmt.Printf("%v and the type is given by %T\n", indexed, indexed)

	for i,v := range str5{
		fmt.Printf("The index is : %v and the value is(rune) it's utf-8 value and not the character : %v and memory %v\n", i, v, &v)
	}
	// since string are nothing but character array they have memory address for each charcater(rune) increment by 4 bits 

	// GO represent strings in your system by using utf-8 instead of ASCII (7 bits)

	// Tip: Easy way to deal with strings is to cast them into array of rune before iteration 
	var str6 = []rune("resumé")
	fmt.Println(str6)

	for i  := range(str6){
		fmt.Printf("The (utf-8) value at index %v is %v \n", i, str6[i])
	}

	// Strings are immutable in go (will cause an compilation error)
	// str5[3] = 'g'

	// Read about Go String Builder : Faster way to concatenate strings
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
