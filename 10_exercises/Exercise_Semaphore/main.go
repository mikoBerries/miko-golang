//semaphore channel using separate boolean in channel like a flag when it's all true means it's done
package main

import "fmt"

func main() {
	c := make(chan int)     //channel int
	done := make(chan bool) //done as flag boolean
	n := 10
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			//will waiting n time until done has value to extract from channel
			<-done
		}
		//close chan when all flag from chan "done" are extracted
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Done program exit")
}
