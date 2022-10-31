package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {

	log.Println("---1---")
	// https://go.dev/play/p/j-EA6003P0 make this dedlock gone
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)

	a := make(chan int, 1)
	go func() {
		a <- 42
	}()
	fmt.Println(<-a)

	log.Println("---2---")
	//https://go.dev/play/p/oB-p3KMiH6

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	//https://play.golang.org/p/_DBRueImEq

	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)

	log.Println("---3---")
	//https://play.golang.org/p/sfyu4Is3FGc :=
	cd := gen()
	receive(cd)

	fmt.Println("about to exit")

	log.Println("---4---")
	// Starting with this code, pull the values off the channel using a select statement
	//https://go.dev/play/p/MvL6uamrJP

	q := make(chan int)
	chn := gen2(q)

	receive2(chn, q)

	fmt.Println("about to exit")

	log.Println("---5---")
	//Show the comma ok idiom starting with this code.
	//https://go.dev/play/p/YHOMV9NYKK

	cnd := make(chan int)
	go func() {
		cnd <- 29
	}()

	v, ok := <-cnd
	fmt.Println(v, ok)

	close(cnd)

	v, ok = <-cnd
	fmt.Println(v, ok) //if there is value ok will send true

	//write a program that
	// ○ puts 100 numbers to a channel
	// ○ pull the numbers off the channel and print them

	channel100 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {

			channel100 <- i
		}
		close(channel100)
	}()

	for v := range channel100 {
		fmt.Println(v)
	}
	log.Println("---6---")

	//● write a program that
	// ○ launches 10 goroutines
	// ■ each goroutine adds 10 numbers to a channel
	// ○ pull the numbers off the channel and print them

	tempChannel := make(chan int)

	var wg sync.WaitGroup
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				for x := 1; x <= 10; x++ {
					tempChannel <- x
				}
				wg.Done()

			}()
		}

		wg.Wait()
		close(tempChannel)
	}()

	for v := range tempChannel {
		fmt.Println(v)
	}

	log.Println("---7---")
}

func receive2(chn <-chan int, q chan int) {
	for {
		select {
		case v := <-chn:
			fmt.Println("v:", v)
		case <-q:
			return

		}
	}
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 0
		close(c)
	}()

	return c
}

func receive(cd <-chan int) {

	for v := range cd {
		fmt.Println("value from channel:", v)

	}
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
