package main

// incase of confusion or for quick summary : alex mux learn go full tutorial: 40 mins mark

import "fmt"
import "time"
import "sync"

var wg = sync.WaitGroup{}
var m = sync.Mutex{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

// storing the result from db call
var res = []string{}

func main(){
	// Go Routines
	fmt.Println("Go Routines Tutorials : Concurrency (Launcing multiple func and have them execute concurrently)")
	fmt.Println("Concurrency is not parallelism : Go can use parallelism (using mutiple cpu cores) to achieve concurrency")
	fmt.Println("Other way takin mutiple tasks and dividing them into subtaks them context switching between substask of different tasks")

	t0 := time.Now()
	for i := 0; i<len(dbData); i++{
		// dbCall(i)
		// to make this func run concurrently use the keyword go before the func call
		// go dbCall(i)
		// if we just do that the program put it in bg and moves to next line....hence no output, so we have to wait for res using sync and waitGroup
		// Wait groups are like counter , add it before the go concurrent fn and in the fn body put done at the end

		wg.Add(1)
		go dbCall(i)
	}
	// wait() to make the counter go back to zero as we added 1 before, implying all the task as has been completed and the rest of code will execute
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v \n", time.Since(t0))

	// before routine the time was : 6.9315485s , after it : 1.6914166s

	// using goroutines to store the data from db and printing it
	fmt.Println("The res from db are: " , res)
	//when we have mutiple threads modifying the same memory location in same time will lead to unexpected result: for eg corrupt memory and  it can lead to missing data , so we shouldn't write like this, instead use mutex to preserve integrity making only one thread to work on the memory block at a time ; m.lock() and m.unlock()



	// ------------------------------ Channels : The main fn for channels is to hold data, thread safety (no race condition: r/w conflict condition) and Listen to Data(when data is added or removed and block code execution till any of it occurs) the datagoroutines communicates data usin channels

	// var c = make(chan int)
	// // channels are like vectors (except they don't adjust size) c : [1]
	// c <- 1
	// var i = <- c
	// // gives a deadlock error  : the code blocks until someone reads from it
	// fmt.Println(i)
	// now channel c is empty and i holds the value 1 , c : []

	// So this isn't the way to use channels use slices/array for storing purposes
	
	
	//---------------------------- Channels are used in conjuction with goRoutines

	// var ch = make(chan int)
	// go process2(ch)
	// // <-ch : directly poping out the value to print rather than storing it in a varaible : the channel is empty after this
	// fmt.Println(<-ch)

	// using channel to get mutiple values using for loop
    var ch = make(chan int)
	go process3(ch)
	for i:= range ch{
		fmt.Println(i)
	}


	// Buffer channel --- Can store mutiple values , ch : [1,2,4,7,9]
	var c = make(chan int, 5)
	go process4(c)

	for i:= range c{
		fmt.Println(i)
		// Some Work
		time.Sleep(time.Second+1)
	}

	// When we use normal channel , the process fn remains open/in memeory, till the main fn is done with the channel in the above loop
	// The process fn has no need to stay active, it can complete it's execution and leave , we use buffer channel for this (extra memory management)

}

func dbCall(i int){
	// simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock()
	// concurrency matters on where we put our lock , because for a particular thread this part is locked , the next running parallely(on other cpu core) can't access this until this is released ....  but if we put the lock after delay declaration , it will make our program non-concurrent , it is due to the fact that the other thread can't make changes to res arr as they can't make no fn call, so they will for lock to be released then they will make fn call.....serially (no-concurrency) : There are other mutex as well RWMutex for read/write locks .... whole db concurrency concepts exact and some OS concepts with parallel threading and data items mutex 
	res = append(res, dbData[i])
	m.Unlock()
	wg.Done()
}

// func process(channel chan int){
// 	channel <- 123
// }

// func process2(channel chan int){
// 	channel <- 123
//   
// }

func process3(ch chan int){
	defer close(ch)
	for i:= 0; i<5; i++{
		ch <- i;
	}
}

func process4(c chan int){
	// it closes the channel just after the execution of this fn :- prevent deadlock error
	defer close(c)
	for i:=0; i<5; i++{
		c <- i;
	}
	fmt.Println("Exiting Process from Process Function")
}