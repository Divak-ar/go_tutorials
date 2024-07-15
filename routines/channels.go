package main

import (
	"fmt";
	"math/rand";
	"time";
)

var Max_Food_Price float32 = 5

func main(){
	var fChannel = make(chan string)
	var websites = []string{"zomato.com","swiggy.com","amazon.com","jeffBezosFucks.com"}

	// we are creating a food search filter which takes max price 5 here, search it concurrently (goRoutines) on websites , if less than max price we return a msg of founded deal on said website

	for i:= range websites{
		// calling go routines 
		go checkFoodPrice(websites[i], fChannel)
	}
	sendMessage(fChannel)
}

func checkFoodPrice(website string, fChannel chan string){
	for{
		time.Sleep(time.Second+1)
		// giving random price to that food items for every website
		var foodPrice = rand.Float32()*30
		if foodPrice <= Max_Food_Price{
			fChannel <- website
			break;
			// the first website with lower foodprice than max will get selected and then we return
		}
	}
}

// read Select statement in channels (basically if-else statement/switch statement) for channels

func sendMessage(fChannel chan string){
	fmt.Printf("\nFound deal on food item at %s", <-fChannel)
}