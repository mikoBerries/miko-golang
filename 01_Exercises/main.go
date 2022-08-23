package main

import (
	"fmt"
)

func main() {
	HandsOnExcise1()
	HandsOnExcise2()
	HandsOnExcise3()
	HandsOnExcise4()
	HandsOnExcise5()
}

func HandsOnExcise5() {
	/*
		1. at the package level scope, using the “var” keyword, create a VARIABLE with the
			IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom
			TYPE “x”
		2. n func main
			a. this should already be done
				i. print out the value of the variable “x”
				ii. print out the type of the variable “x”
				iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
				iv. print out the value of the variable “x”
			b. now do this
				1.now use CONVERSION to convert the TYPE of the VALUE stored in “x”
				to the UNDERLYING TYPE
					1. then use the “=” operator to ASSIGN that value to “y”
					2. print out the value stored in “y”
					3. print out the type of “y”
	*/
	type hotdog int
	var x hotdog
	var y int
	x = 42
	y = int(x)

	fmt.Printf("%v \t %T \n", y, y) //still int
	fmt.Printf("%v \t %T \n", x, x)
}

func HandsOnExcise4() {
	/*
		Hands-on exercise #4
		● FYI - nice documentation and new terminology “underlying type”
		○ https://golang.org/ref/spec#Types
		For this exercise
		1. Create your own type. Have the underlying type be an int.
		2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
		3. in func main
			a. print out the value of the variable “x”
			b. print out the type of the variable “x”
			c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
			d. print out the value of the variable “x”
	*/

	fmt.Println("HandsOnExcise4 ->Start")

	type sleppyCat int
	var x sleppyCat
	fmt.Printf("%v \t %T \n", x, x)
	x = 42
	fmt.Printf("%v \t %T \n", x, x)

	fmt.Println("HandsOnExcise4 ->Start")
}

func HandsOnExcise3() {
	/*
		Hands-on exercise #3
		Using the code from the previous exercise,
		1. At the package level scope, assign the following values to the three variables
			a. for x assign 42
			b. for y assign “James Bond”
			c. for z assign true
		2. in func main
			a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
			returned value of TYPE string using the short declaration operator to a
			VARIABLE with the IDENTIFIER “s”
			b. print out the value stored by variable “s”
	*/

	fmt.Println("HandsOnExcise3 ->Start")

	x, y, z := 42, "James Bond", true
	// s := fmt.Sprintf("%d %s %t", x, y, z)
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
	fmt.Println("HandsOnExcise3 ->END")

}

func HandsOnExcise2() {
	/*
		Hands-on exercise #2
		1. Use var to DECLARE three VARIABLES. The variables should have package level
		scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
		variables and make sure the variables are of the following TYPE (meaning they can
		store VALUES of that TYPE).
			a. identifier “x” type int
			b. identifier “y” type string
			c. identifier “z” type bool
		2. in func main
			a. print out the values for each identifier
			b. The compiler assigned values to the variables. What are these values called?
	*/
	fmt.Println("HandsOnExcise2 ->Start")

	var x int
	var y string
	var z bool
	fmt.Print(x)
	fmt.Printf("\t %T\n", x)
	fmt.Print(y)
	fmt.Printf("\t %T\n", y)
	fmt.Print(z)
	fmt.Printf("\t %T\n", z)

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
