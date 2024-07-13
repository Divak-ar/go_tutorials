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

func main() {

	// Struct and Interfaces in GO
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 12}
	// var myEngine gasEngine = gasEngine{25,12}
	// myEngine.mpg = 53 
	fmt.Println(myEngine)
	fmt.Println(myEngine.mpg)

	var myEngine2 gasEngine = gasEngine{25, 12, owner{"Adler"}}
	fmt.Println(myEngine2.mpg, " ", myEngine2.gallons, " ", myEngine2.ownerInfo.name)



}