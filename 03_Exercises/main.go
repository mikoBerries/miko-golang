package main

import "fmt"

func main() {
	/*Hands-on exercise #1
	Print every number from 1 to 10,000*/
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
	/*Hands-on exercise #2
	Print every rune code point of the uppercase alphabet three times. Your output should look like
	this:
	65
	U+0041 'A'
	U+0041 'A'
	U+0041 'A'
	66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
	… through the rest of the alphabet characters
	'ascii
	*/
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
	/*Hands-on exercise #3
	Create a for loop using this syntax
	● for condition { }
	Have it print out the years you have been alive.

	Hands-on exercise #8
	Create a program that uses a switch statement with no switch expression specified.

	*/
	birthDate := 1994
Exit:
	for {

		switch {
		case birthDate == 2022:
			break Exit
		default:
			fmt.Println(birthDate)
			birthDate++
		}
	}

	/*Hands-on exercise #5
	Print out the remainder (modulus) which is found for each number between 10 and 100 when it
	is divided by 4.
	*/
	for i := 10; i <= 100; i++ {
		fmt.Println(i, "% 4 =", i%4)
	}

	/*Hands-on exercise #6
	Create a program that shows the “if statement” in action.
	Hands-on exercise #7
	Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

	*/
	if false {
		fmt.Println("something 1")
	} else if !true {
		fmt.Println("something 2")
	} else {
		fmt.Println("something 3")
	}
	/*Hands-on exercise #9
	Create a program that uses a switch statement with the switch expression specified as a
	variable of TYPE string with the IDENTIFIER “favSport”.
	*/
	favsport := "fotball~"
	switch favsport {
	case "fotball~":
		fmt.Println("favsport is fotball~")
	case "other than fotball~":
		fmt.Println("favsport is other than fotball~")
	case "aanother things ~":
		fmt.Println("favsport is aanother things ~")
	default:
		fmt.Println("favsport other things ")
	}
	/*Hands-on exercise #10
	Write down what these print:
	● fmt.Println(true && true) -->true
	● fmt.Println(true && false) -->false
	● fmt.Println(true || true)--> true
	● fmt.Println(true || false) -->true
	● fmt.Println(!true) -->false
	*/
}
