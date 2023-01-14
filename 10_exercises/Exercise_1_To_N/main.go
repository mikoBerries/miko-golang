//one channel to many reciever chan
package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	//inserting many data to channel "c"
	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	//recieving "n" times and flagging each go routine using "done" channel
	for i := 0; i < n; i++ {
		go func() {
			for value := range c {
				fmt.Println(value)
			}
			done <- true
		}()
	}

	//recieving "n" times from "done" channel
	for i := 0; i < n; i++ {
		<-done
	}
	fmt.Println("Done - program exit")
}
