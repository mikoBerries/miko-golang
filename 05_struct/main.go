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
	//A struct is an composite data type. (composite data types, aka, aggregate data types, aka, complex data types).
	p1 := person{
		first: "james",
		last:  "bond",
		age:   10,
	}
	p2 := person{
		first: "miss",
		last:  "moneyponey",
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)

	s1 := secretAgent{
		p1, true,
	}
	s2 := secretAgent{
		person: person{
			first: "james",
			last:  "bondong",
		},
		lol: true,
	}

	fmt.Println(s1)
	fmt.Println(s2)
	//innner struct value promoted to outer like inheritance

	fmt.Println(s1.first, s1.last, s1.lol)

	//Anonymous structs making struct without declaring in type
	p45 := struct {
		first, last string
		ages        int
	}{
		first: "first name",
		last:  "last name",
		ages:  10,
	}
	fmt.Println(p45)
}
