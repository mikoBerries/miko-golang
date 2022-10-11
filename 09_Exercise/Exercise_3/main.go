package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int64
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	/*
				Hands-on exercise #3
			● Using goroutines, create an incrementer program
				○ have a variable to hold the incrementer value
				○ launch a bunch of goroutines
			■ each goroutine should
			● read the incrementer value
				○ store it in a new variable
		● yield the processor with runtime.Gosched()
			● increment the new variable
			● write the value in the new variable back to the incrementer
			variable
			● use waitgroups to wait for all of your goroutines to finish
			● the above will create a race condition.
			● Prove that it is a race condition by using the -race flag
	*/
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go incrementer()
	}
	wg.Wait()
}

func incrementer() {
	defer wg.Done()
	mutex.Lock()
	counter += 1
	runtime.Gosched()
	fmt.Println(counter)
	mutex.Unlock()
}
