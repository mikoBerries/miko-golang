package main

import (
	"fmt"
	"runtime"
)

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
}
