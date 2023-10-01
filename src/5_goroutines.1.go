package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// In this script we'll review how to approach concurrency with goroutines
// - goroutines are lightweight threads
// - goroutines are run with the 'go' keyword and a worker function
// - goroutines are scheduled by the go runtime
// - the main thread will not wait for goroutines to finish unless you use a waitgroup
// - the main thread is a goroutine itself
// for now we'll focus on this example

func worker(id int) {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	fmt.Printf("Worker #%d started\n", id)
	seconds := rand.Intn(5) + 1                // generate a random number between 1 and 5
	wt := time.Duration(seconds) * time.Second // convert the number to a time.Duration
	time.Sleep(wt)                             // sleep the thread for the duration
	fmt.Printf("Worker #%d finished\n", id)
}

func main() {
	fmt.Println("Program started")
	// Failing example:
	// to create a goroutine we'll use the 'go' keyword
	// go worker(1) // this will run the worker in a separate thread
	// lets create 10 goroutines
	// for i := 0; i < 10; i++ {
	// 	go worker(i)
	// }
	// fmt.Println("Program finished")
	// in this example the main program thread will finish before the goroutines
	// since they are running in different threads
	// only "Program started" and "Program finished" will be printed
	// to fix this we can use a waitgroup as follows

	wg := sync.WaitGroup{} // a way to wait for goroutines to finish
	// lets start 10 goroutines with waitgroups
	for i := 0; i < 10; i++ {
		wg.Add(1)         // add a goroutine to wait for
		go func(id int) { // closure with access to the waitgroup
			// defer makes wg.Done() run at the end of the function
			defer wg.Done() // mark the goroutine as finished
			worker(id)
		}(i)
	}

	wg.Wait() // wait for all goroutines to finish, blocks the main thread
	fmt.Println("Program finished")
}
