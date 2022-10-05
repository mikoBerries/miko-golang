package main

import "fmt"

func main() {
	//normal loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//nested loop
		for x := 0; x < i; x++ {
			if x%2 == 0 {

				fmt.Printf("\t%d", x)
			}
		}
	}

	a := 0
	for { //eternal loop
		if a > 10 {
			break
		} else if a == 5 {
			continue
		}
		a++
		fmt.Println(a)
	}

}
