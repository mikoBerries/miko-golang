package main

import "fmt"

type person struct {
	first string
	last  string
	age   int //default value is not set is 0
}

//embed struct
type secretAgent struct {
	person
	lol bool
}

//interface
type human interface {
	speak()
	//all who has speak func is a humnan (polymorph)
	//empty interface func in meaning every one are a human
}

type hotdog int

func main() {
	bar("something")
	defer fmt.Println("1st defer")
	if v, ok := barbar("my word"); ok {
		fmt.Println(v)
	}
	fmt.Println(boo(1, 2, 3, 4, 5, 6, 7))
	xi := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println(boo(xi...))
	sa1 := secretAgent{
		person{"james", "bondong", 10},
		true,
	}
	sa1.speak()

	defer fmt.Println("2st defer")
	//defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns
	// multiple defer statements are allowed in the same program and they are executed in LIFO(Last-In, First-Out)/stack order as shown in Example 2.

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	//polymorph person an secret agent are human
	fobar(p1)
	fobar(sa1)

	// conversion from type hotdog become int
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	//Anonymous func declaration and call
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("Anonymous func ran with param : ", x)
	}(42) //insert param like making Anonymous struct

	//func expresion assign func() to variable
	f := func() {
		fmt.Println("my first func expression")
	}
	f2 := func(s string) {
		fmt.Println("func expression with param", s)
	}
	f()
	f2("this is mya string")

	barbarbar()(2022)

	//passing function as param
	math(23, 7, add)
	math(23, 7, multiply)

	//closure
	a := incrementor()
	b := incrementor()
	fmt.Println(a()) //1
	fmt.Println(a()) //2 not reseting because var X outside of clouser braket =>{}
	fmt.Println(a()) //3
	fmt.Println(a()) //4
	fmt.Println(b())
	fmt.Println(b())

	//Recursion is bad for memory use loop or other else LOL
	mytotal := func(n int) int {
		total := 1
		for ; n > 0; n-- {
			total *= n
		}
		return total
	}(4)
	fmt.Println(mytotal)
}

//everything in go is pass by value
func bar(s string) {
	fmt.Println(s)
}

func barbar(st string) (string, bool) {
	return fmt.Sprintf("Hello from other" + st), true
}

func boo(x ...int) string {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return fmt.Sprintf("total is %d", sum)
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

//Polymorphism is the ability of a message to be displayed in more than one form.
// Polymorphism is considered as one of the important features of Object-Oriented Programming
// and can be achieved during either at runtime or compile time.
// Golang is a light-Object Oriented language and supports polymorphism through interfaces only.
//https://www.geeksforgeeks.org/pointer-to-a-struct-in-golang/?ref=lbp
func fobar(h human) { //fobar func accept h as human

	fmt.Println("I was passed into bar", h)
	fmt.Printf(">>>%T \n", h)
	switch h.(type) { //asertion (type case)
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
}

//func returning func() as type
func barbarbar() func(number int) string {
	return func(number int) string {
		return fmt.Sprintf("this is :", number)
	}
}

func add(x, y int) {
	fmt.Println(x + y)
}

func multiply(x, y int) {
	fmt.Println(x * y)
}

func math(x int, y int, f func(a, b int)) {
	f(x, y)
}

//clouor variable
func incrementor() func() int {
	var x int           //var x outside thre returning of func
	return func() int { //returning func
		x++
		return x
	}
}
