package main

import "fmt"

func main() {
	bar("something")
	if v, ok := barbar("my word"); ok {
		fmt.Println(v)
	}
	fmt.Println(boo(1, 2, 3, 4, 5, 6, 7))
	xi := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println(boo(xi...))

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
