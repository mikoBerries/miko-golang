package main

import "fmt"

//global var
var x string // init value string -> ""
var y int    // init value int -> 0

type hotdog int // similar to java class declaration
var sausage hotdog

func main() {
	fmt.Println("123") //1 line

	fmt.Printf("123\n") //with costume format
	fmt.Printf("%d - %b \n", 42, 42)

	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	// 42 (decimal)   101010 (binary)    0X2A(Hexa)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	fmt.Print("123")

	//looping
	for i := 0; i <= 100; i++ {
		//:= is known as the short declaration operator
		// use " _ " to dumping some unused throwback val
		n, e := fmt.Println(i)
		fmt.Println(n, e)
	}
	dothings()
	doOtherThings()

}

func doOtherThings() {
	sausage = 10
	fmt.Println("sausage value : ", sausage)   //10
	fmt.Printf("sausage type :%T \n", sausage) //main.hotdog (package.class/type)
}
func dothings() {
	fmt.Printf("value x is :" + x) // ""
	fmt.Printf("%T \n", x)         // int
	fmt.Print("value y is :")      // 0
	fmt.Printf("%T \n", y)         // string

	fmtedString := fmt.Sprintf("there is %d wolf in the shed number %d", 10, 1)
	fmt.Println(fmtedString)
}
