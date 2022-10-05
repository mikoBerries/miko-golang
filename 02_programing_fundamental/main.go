package main

import (
	"fmt"
	"runtime"
)

const x = "new string" //automated gengerat string type
const xy string = "new string"
const (
	j = iota
	k
	l
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
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
	const p = "death & taxes"

	const q = 42
	fmt.Println("x - ", x)
	fmt.Println("xy - ", xy)
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	// a CONSTANT is a simple unchanging value

	//iota automated increment interger for constant

	fmt.Println(">> IOTA <<")
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)

	//bit shifting using  << or >>

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
