package main

import (
	"fmt"
)

const (
	x        = 10
	y string = "this is typed constant"
)

const (
	year1 = 2022 + iota
	year2 = 2022 + iota
	year3 = 2022 + iota
	year4 = 2022 + iota
)

func main() {
	//Hands-on exercise #1
	//Write a program that prints a number in decimal, binary, and hex
	x := 10
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#X\n", x)

	/*Hands-on exercise #2
	Using the following operators, write expressions and assign their values to variables
	g. ==
	h. <=
	i. >=
	j. !=
	k. <
	l. >
	*/
	g := (42 == 43)
	h := (42 <= 43)
	i := (42 >= 43)
	j := (42 != 43)
	k := (42 < 43)
	l := (42 > 43)
	fmt.Println(g, h, i, j, k, l)

	/*Hands-on exercise #2
	Create TYPED and UNTYPED constants. Print the values of the constants*/
	fmt.Println(x)
	fmt.Println(y)

	/*Hands-on exercise #4
	Write a program that
	● assigns an int to a variable
	● prints that int in decimal, binary, and hex
	● shifts the bits of that int over 1 position to the left, and assigns that to a variable
	● prints that variable in decimal, binary, and hex
	*/

	myint := 100

	fmt.Printf("%d\n", myint)  //dec
	fmt.Printf("%b\n", myint)  //bin
	fmt.Printf("%#X\n", myint) //hex

	myint = myint << 1

	fmt.Printf("\n%d\n", myint) //dec
	fmt.Printf("%b\n", myint)   //bin
	fmt.Printf("%#X\n", myint)  //hex

	/* Hands-on exercise #5
	Create a variable of type string using a raw string literal. Print it
	*/

	mystring := `here is something
	as 
	a 
	raw string
	literal
	"you see"
	another thing`
	fmt.Println(mystring)

	/*Hands-on exercise #6
	Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
	*/

	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
