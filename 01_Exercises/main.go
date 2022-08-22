package main

import "fmt"

func main() {
	HandsOnExcise1()
	HandsOnExcise2()

}

func HandsOnExcise2() {
	fmt.Println("HandsOnExcise2 ->Start")

	fmt.Println("HandsOnExcise2 ->END")
}

func HandsOnExcise1() {
	// Hands-on exercise #1
	// 1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the
	// IDENTIFIERS “x” and “y” and “z”
	// a. 42
	// b. “James Bond”
	// c. true
	// 2. Now print the values stored in those variables using
	// a. a single print statement
	// b. multiple print statements
	fmt.Println("HandsOnExcise1 ->Start")
	x, y, z := 42, "James Bond", true
	fmt.Print(x, "\n")
	fmt.Print(y, "\n")
	fmt.Print(z, "\n")
	fmt.Println(x, y, z)

	fmt.Println("HandsOnExcise1 ->END")
}
