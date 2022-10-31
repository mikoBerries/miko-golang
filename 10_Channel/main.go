package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//buffered and unbuffered channel
	//unbuffered channel is used to perform synchronous communication between goroutines
	//buffered channel are is used for perform asynchronous communication.
	unBufferedExample()
	fmt.Println()
	bufferedExample()
	log.Println("1")
	//directional channell
	// chanR := make(<-chan int) //chanR is recieve value only (obtain from chan) <-chanR
	// chanS := make(chan<- int) //chanS is send value only (inserting to chan) chanS <- 10
	// <-chanR
	// chanS <- 10
	log.Println("2")

	myChannel := make(chan int)
	//using directional channel
	go sendOnlyChannel(myChannel)
	recieveOnlyChannel(myChannel) //locked until channel get value
	log.Println("3")

	//Select statement to pull any channel that ready to pull value
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send function
	go send(eve, odd, quit)
	receive(eve, odd, quit)

	log.Println("4")

	chann1 := make(chan int)
	chann2 := make(chan int)
	go populate(chann1) //filling channel

	go fanOutIn(chann1, chann2)

	for v := range chann2 {
		fmt.Println(v)
	}

	log.Println("5-- fan in out")

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel() // cancel when we are finished to close all go routine if program suddenly stop

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	log.Println("5-- context for go routine to prevent leaking")
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // returning not to leak the goroutine
			case dst <- n:
				fmt.Println("n:", n)
				n++
			}
		}
	}()
	return dst
}

func fanOutIn(chann1 chan int, chann2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ { //launch 10 go routines
		go func() {
			for v := range chann1 { //taking value from chnn1
				func(v2 int) {
					chann2 <- timeConsumingWork(v2) //do work from that value
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(chann2)
}
func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
func populate(chann1 chan int) {
	for i := 0; i < 100; i++ {
		chann1 <- i
	}

	close(chann1)
}

func send(e chan<- int, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- true
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case v := <-quit:
			fmt.Println("the value received from the quit channel:", v)
			return
		}
	}
}

func recieveOnlyChannel(myChannel <-chan int) {
	fmt.Println(<-myChannel)

}

func sendOnlyChannel(myChannel chan<- int) {
	myChannel <- 10000
}

func unBufferedExample() {
	//using un-buffered channel because
	//the channel are perform sends (c<-values) and receives (<-c) value synchronous
	//diffrent go routines are sending and recieving from channel.
	c := make(chan int)
	go func() {
		c <- 10
		c <- 20
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func bufferedExample() {
	//using buffered channel because
	//the channel are perform sends (c<-values) and receives (<-c) value asynchronous.
	c := make(chan int, 5)
	c <- 10
	c <- 20
	c <- 30
	c <- 40
	c <- 50

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
