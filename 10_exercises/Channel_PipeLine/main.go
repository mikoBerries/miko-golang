//using channel like a pipeline from func to func that have chan param
package main

import "fmt"

func main() {

	data := generate()
	factorial := factorialMath(data)

	for value := range factorial {
		fmt.Println(value)
	}

}
func generate() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out) //close channel after done generating
	}()

	return out
}

func factorialMath(data <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for value := range data {
			out <- fact(value)
		}
		close(out)
	}()
	return out
}

//calculate factorial function
func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
