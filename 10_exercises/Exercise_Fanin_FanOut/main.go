//Simple Usage of Fan in Fan out channel model in golang
package main

import (
	"fmt"
	"sync"
)

func main() {
	//source data
	dataInput := generate(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 17, 18, 19, 20)

	//fan out 2 channel or more to faster processing data
	ch1 := sq(dataInput)
	ch2 := sq(dataInput)

	//fan in from 2 or more channel (merged it as 1 channel)
	for value := range mergeChannel(ch1, ch2) {
		fmt.Println(value)
	}
}

func generate(numbers ...int) chan int {
	output := make(chan int)
	go func() {
		for _, value := range numbers {
			output <- value
		}
		//close generate channel
		close(output)
	}()
	return output
}

func sq(data chan int) chan int {
	output := make(chan int)
	go func() {
		for value := range data {
			output <- value * value
		}
		close(output)
	}()
	return output
}

func mergeChannel(channels ...chan int) chan int {
	mergedChannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels)) //add waitgroup as length channels data
	//merge from ...chan int to "mergedChannel"
	for _, c := range channels {
		go func(ch chan int) {
			for value := range ch {
				mergedChannel <- value
			}
			wg.Done()
		}(c)
	}
	//waiting wg all done then close channel
	go func() {
		wg.Wait()
		close(mergedChannel)
	}()

	return mergedChannel
}
