package main

import "fmt"

type person struct {
	first, last string
}

func main() {
	/*
		Hands-on exercise #1
		● Create a value and assign it to a variable.
		● Print the address of that value.s
	*/
	myvalue := 100
	fmt.Println(&myvalue)

	/*
		Hands-on exercise #2
		● create a person struct
		● create a func called “changeMe” with a *person as a parameter
			○ change the value stored at the *person address
		■ important
		● to dereference a struct, use (*value).field
			○ p1.first
			○ (*p1).first
		● “As an exception, if the type of x is a named pointer type and (*x).f
		is a valid selector expression denoting a field (but not a method),
		x.f is shorthand for (*x).f.”
			○ https://golang.org/ref/spec#Selectors
		● create a value of type person
			○ print out the value
		● in func main
			○ call “changeMe”
		● in func main
			○ print out the value
	*/
	p1 := person{
		"first", "last",
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	changeme(&p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
}

func changeme(p *person) {
	(*p).first = "changed first"
	p.last = "change last"
}
