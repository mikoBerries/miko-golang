package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go doSomething("this is go routine one")
	go doSomething("this is go routine two")
	wg.Wait()
}

func doSomething(s string) {
	defer wg.Done()
	fmt.Println(s)
}
