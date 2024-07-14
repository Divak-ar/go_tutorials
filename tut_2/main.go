package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	ownerInfo owner
}

type owner struct{
	name string
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

// methods : these are fns that are tied to struct and have access to struct instance itself

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh*e.kwh
}

// Interfaces : 

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("You can make it there!")
	}else{
		fmt.Println("Need to fuel up first!")
	}
}



func main() {

	// Struct and Interfaces in GO

	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 12}
	// var myEngine gasEngine = gasEngine{25,12}
	// myEngine.mpg = 53 
	fmt.Println(myEngine)
	fmt.Println(myEngine.mpg) 

	var myEngine2 gasEngine = gasEngine{25, 12, owner{"Adler"}}
	fmt.Println(myEngine2.mpg, " ", myEngine2.gallons, " ", myEngine2.ownerInfo.name)


	// anonymous structs : they can be define by below syntax but can't be reused ... secondly when defining we have to assign values
	var myEng3 = struct{
		mpg uint8
		gallons uint8
		car string
	}{23,5,"BMW"}

	fmt.Println("Can only be initialized once" , myEng3.mpg, myEng3.gallons)

	myEng3.mpg = 45;
	fmt.Println(myEng3.mpg," ", myEng3.car)

	// we can't create another strcut that has mpg and gallons and cars ..... for mutiple use don't use anonymous structs

	// Methods : Adding a method to a struct in Golang is simple. We define the method outside of the struct definition and use the struct as the receiver for the method. The receiver is a parameter that provides access to the fields of the struct and perform some operation on it.

	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())
	fmt.Println()
	var eEngine1 electricEngine = electricEngine{23,4}
	fmt.Printf("Total miles left in tank: %v \n", eEngine1.milesLeft())

	// what if we want to have a fn that takes any type of engine instead of specific struct and methods like gas and electric engine : Interfaces comes to play
	// for eg : func canMakeIt(e gasEngine, miles uint8){
		// if miles <= e.milesLeft(){
		// 	fmt.Println("You can make it there!")
		// }else{
		// 	fmt.Println("Need to fuel up first!")
		// }
	// } this func is generalized for gasEngine ..... If we use interface we can make it avaiable for any type of engine
	// by replacing gasEngine with interface engine with method milesLeft(), it signifies it can take any object that has milesLeft() method with it

	canMakeIt(myEngine, 52)
	canMakeIt(eEngine1, 92)


	// ------------------------------------

	// Pointers ; they are variables that stores the location of some data (*)

	var p *int32
	fmt.Println(p)
	// doesn't point to anything now : nil
	var ptr *int32 = new(int32)
	fmt.Println("This Points to any int32 address in memory: Specially allocated to it ", ptr)
	fmt.Println("Value stored at this location is: {default value for int} ", *ptr)

	*ptr = 10
	fmt.Println("Value assigned is : " , ptr)


	var v int32 = 45
	p = &v

	fmt.Printf("ptr p now ref to mem address of v : %v and the value of v is : %v \n", p, *p)


	// Slices uses pointer under the hood , so if we copy the value of one slice to another , it will be reference , so any changes made into one will reflect the other 

	var slice = []int32{2,5,7}
	var sliceCopy = slice
	sliceCopy[2] = -1
	fmt.Println(sliceCopy)
	fmt.Println(slice)

	// Pointers in func

	var a1 = [3]float32{2,3,4}
	fmt.Printf("\nThe memory location of the a1 is: %p \n", &a1)
	var a2 [3]float32 = square(a1)
	fmt.Println("The modified val is: ", a2)

	// passing by address and pointers
	fmt.Printf("\nThe memory location of the a1 is: %p \n", &a1)
	var a3 [3]float32 = addressSquare(&a1)
	fmt.Println("The modified val is: ", a1)
	fmt.Println("The modified val is: ", a3)
	
	// Since we passed by reference(address) the value of a changes as well
}

func square(a [3]float32) [3]float32{
	fmt.Printf("The memory location of the a {int func body} is: %p \n", &a)
	// the memory location is different cause it is pass by value so a1 doesn't changes it returns modified value to a2
	// we can use pointers so that we don't create another array here a to store a1 and return it (pass address and use pointer in func)
	for i := range a{
		a[i] = a[i]*a[i]
	}

	return a
}

func addressSquare(a *[3]float32) [3]float32{
	fmt.Printf("The memory location of the a {int func body: using poiters} is: %p \n", a)
	for i := range a{
		a[i] = a[i]*a[i]
	}

	return *a
}
