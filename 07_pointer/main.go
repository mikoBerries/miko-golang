package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220

	var b = &a
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value\

	//usage of pointer to change value witoud passing to much data just addres &value
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
	//method set
	//Receivers Values
	// -----------------------------------------------
	// (t T) T and *T   setting not pointer type of receivers can cake type or pointer to that type
	// (t *T) *T		receiver  just accept pointer of initiaed value in method
}

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
