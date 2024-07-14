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


}

func dbCall(i int){
	// simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock()
	// concurrency matters on where we put our lock , because for a particular thread this part is locked , the next running parallely(on other cpu core) can't access this until this is released ....  but if we put the lock after delay declaration , it will make our program non-concurrent , it is due to the fact that the other thread can't make changes to res arr as they can't make no fn call, so they will for lock to be released then they will make fn call.....serially (no-concurrency)
	res = append(res, dbData[i])
	m.Unlock()
	wg.Done()
}