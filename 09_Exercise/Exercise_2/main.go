package main

import "fmt"

type person struct {
	First, Last string
}
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.First, p.Last)
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p := person{"agus", "budi"}
	p.speak()
	saySomething(&p)
}
