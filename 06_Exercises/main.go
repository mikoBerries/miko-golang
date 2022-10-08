package main

import (
	"fmt"
	"math"
)

type person struct {
	first, last string
	age         int
}
type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}
type shape interface {
	area() (area float64)
}

func main() {
	/*
		Hands on exercise #1
		○ create a func with the identifier foo that returns an int
		○ create a func with the identifier bar that returns an int and a string
		○ call both funcs
		○ print out their results
	*/
	fmt.Println(foo())
	fmt.Println(bar())

	/*
		Hands-on exercise #2
		● create a func with the identifier foo that
		○ takes in a variadic parameter of type int
		○ pass in a value of type []int into your func (unfurl the []int)
		○ returns the sum of all values of type int passed in
		● create a func with the identifier bar that
		○ takes in a parameter of type []int
		○ returns the sum of all values of type int passed in
	*/

	fmt.Println(foo2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...))
	fmt.Println(bar2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
	/*
		Hands-on exercise #3
		Use the “defer” keyword to show that a deferred func runs after the func containing it exits
	*/

	defer fmt.Println("this is defer") //defer work after fun(){} exit
	fmt.Println("printlin under the defer")

	/*
			Hands-on exercise #4
		● Create a user defined struct with
		○ the identifier “person”
		○ the fields:
		■ first
		■ last
		■ age
		● attach a method to type person with
		○ the identifier “speak”
		○ the method should have the person say their name and age
		● create a value of type person
		● call the method from the value of type person
	*/
	p1 := person{
		"agus", "budi",
		20,
	}
	fmt.Println(p1.speak())
	/*
		Hands-on exercise #5
		● create a type SQUARE
		● create a type CIRCLE
		● attach a method to each that calculates AREA and returns it
		○ circle area= π r 2
		○ square area = L * W
		● create a type SHAPE that defines an interface as anything that has the AREA method
		● create a func INFO which takes type shape and then prints the area
	*/

	fmt.Println(info(square{10, 10}))
	fmt.Println(info((circle{24})))

	myFunc := func(x string) {
		fmt.Println("this is my ", x)

	}
	myFunc("lol")
	/*
		Hands-on exercise #8
		● Create a func which returns a func
		● assign the returned func to a variable
		● call the returned func
	*/
	myreturnfunc := HOE8()
	myreturnfunc("try this one")

	/*
		Hands-on exercise #9
		A “callback” is when we pass a func into a func as an argument. For this exercise,
		● pass a func into a func as an argumen
	*/
	mathfuncA := add()
	mathfuncD := devide()
	calculate(1, 2, mathfuncA)
	calculate(100, 2, mathfuncD)
	/*
		Hands-on exercise #10
		Closure is when we have “enclosed” the scope of a variable in some code block. For this
		hands-on exercise, create a func which “encloses” the scope of a variable:
	*/
	xy := 10
	{
		xy := 100
		fmt.Println("inside", xy)
	}
	fmt.Println("outside", xy)
	f := hoe10()
	f()
	f()
	f()
	f()
	fmt.Println(f())
}

func (p person) speak() (speaking string) {
	speaking = fmt.Sprint("i'm ", p.first, p.last)
	return
}

func foo() (number int) {
	number = 10
	return
}

func bar() (number int, word string) {
	number = 100
	word = "this is word"
	return
}

func foo2(x ...int) (sum int) {
	for _, v := range x {
		sum += v
	}
	return
}

func bar2(x []int) (sum int) {
	for _, v := range x {
		sum += v
	}
	return
}

func (s square) area() (area float64) {
	area = s.length * s.width
	return
}
func (c circle) area() (area float64) {
	area = math.Pi * c.radius * 2
	return
}

func info(s shape) float64 {
	switch s := s.(type) {
	case circle:
		return s.area()
	case square:
		return s.area()
	default:
		return 0
	}
}

func HOE8() func(s string) {
	return func(s string) {
		fmt.Println("this is returned func" + s)
	}

}

func add() func(x, y int) {
	return func(x, y int) {
		fmt.Println("added val :", x+y)
	}
}
func devide() func(x, y int) {
	return func(x, y int) {
		fmt.Println("devide value:", x/y)
	}
}
func calculate(x int, y int, f func(x, y int)) {
	f(x, y)
}
func hoe10() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
