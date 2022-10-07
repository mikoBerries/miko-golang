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
	fmt.Println("Iam ", s.first, s.last)

}

//Polymorphism is the ability of a message to be displayed in more than one form.
// Polymorphism is considered as one of the important features of Object-Oriented Programming
// and can be achieved during either at runtime or compile time.
// Golang is a light-Object Oriented language and supports polymorphism through interfaces only.
//https://www.geeksforgeeks.org/pointer-to-a-struct-in-golang/?ref=lbp
