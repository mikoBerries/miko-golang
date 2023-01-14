//exercise N value channel to 1 channel
package main

import (
	"fmt"
	"sync"
)

var (
	// mutex sync.Mutex
	wg sync.WaitGroup
)

func main() {
	c := make(chan int)
	//N value of go routine channel inserting
	wg.Add(2)
	//first go routine to insert value to chan int
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	//second go routine to insert value to chan int
	go func() {
		for i := 5; i < 20; i++ {
			c <- i
		}
		wg.Done()
	}()

	//go routine to wait all WG done then close channel c
	go func() {
		wg.Wait()
		close(c)
	}()

	//1 for range to extrect value from channel
	for n := range c {
		fmt.Println(n)
	}
	fmt.Println("done program exit")
}
