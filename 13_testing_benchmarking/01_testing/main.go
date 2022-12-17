//package for testting go lang test
package main

import "fmt"

func main() {
	fmt.Println(Mathadding(1, 2))
	fmt.Println(Mathadding(10, 20))
	fmt.Println(Mathadding(30, 0))
}

func Mathadding(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum //pass
	// return sum + 1 //fail
}
