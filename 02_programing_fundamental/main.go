package main

import (
	"fmt"
	"runtime"
)

const a = "string"

func main() {
	//selected type by go
	a := 12
	b := "ABC"
	c := 12.123123

	fmt.Printf("%v\t%T \n", a, a)
	fmt.Printf("%v\t%T \n", b, b)
	fmt.Printf("%v\t%T\n", c, c)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	const p = "death & taxes"

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	// a CONSTANT is a simple unchanging value

}
