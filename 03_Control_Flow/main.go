package main

import "fmt"

func main() {
	//normal loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//nested loop
		for x := 0; x < i; x++ {
			if x%2 == 0 {

				fmt.Printf("\t%d", x)
			}
		}
	}

	a := 0
	for { //eternal loop

		a++
		fmt.Println(a)
		if a > 10 {
			break
		} else if a == 5 {
			continue
		}
	}
	//ascii print
	//%#x hex val
	//%#U unicode
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
	//boolean if
	myBool1 := true
	myBool2 := false
	if myBool1 {
		fmt.Println(myBool1)
	}
	if myBool2 {

		fmt.Println(myBool2)
	} else if !myBool2 {

		fmt.Println(!myBool2)
	}

	if x := 42; x == 42 { //x only in if scope

		fmt.Println(x)
	}
	//fallthrough execute insede next case :
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("default thiungs ")
	}
	fmt.Println()
	//if else expretion
	fmt.Printf("true && true\t %v\n", true && true)
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", true || true)
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)

}
