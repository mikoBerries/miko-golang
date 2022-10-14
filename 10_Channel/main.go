package main

import "fmt"

func main() {
	//buffered and unbuffered channel
	//unbuffered channel is used to perform synchronous communication between goroutines
	//buffered channel are is used for perform asynchronous communication.
	unBufferedExample()
	fmt.Println()
	bufferedExample()

	//directional channell
	chanR := make(<-chan int) //chanR is recieve value only (obtain from chan) <-chanR
	chanS := make(chan<- int) //chanS is send value only (inserting to chan) chanS <- 10
	_ = <-chanR
	chanS <- 10

	myChannel := make(chan int)
	//using directional channel
	go sendOnlyChannel(myChannel)
	recieveOnlyChannel(myChannel) //locked until channel get value
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
