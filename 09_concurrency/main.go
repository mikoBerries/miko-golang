package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg      sync.WaitGroup
	counter int32
	mutex   sync.Mutex
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)       //OS               windows
	fmt.Println("ARCH\t\t", runtime.GOARCH)   //ARCH             amd64
	fmt.Println("CPUs\t\t", runtime.NumCPU()) //CPUs             4
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100

	wg.Add(gs)

	//mutex.WRlock() // write and read lock
	//mutex.Rlock() // read lock
	//see at go doc

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock() //lock this block of code
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mutex.Unlock() //unlock
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
	concurrency1()
	// time.Sleep(1*time.Second)
	fmt.Println("CPUs\t\t", runtime.NumCPU())             //CPUs             4
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine()) //			2

}

func concurrency1() {
	/*
		Goroutines
			Concurrency in Golang is the ability for functions to run independent of each other.
			Goroutines are functions that are run concurrently.
			Golang provides Goroutines as a way to handle operations concurrently.
	*/

	data := make(chan int) // Declare a unbuffered channel
	defer close(data)      //close channel
	s := []string{"https://www.golangprograms.com"}
	for _, v := range s {
		wg.Add(1)
		fmt.Printf("wg: %v\n", wg)
		go responseSize(v, data)

	}
	fmt.Println("done forr now -> waiting wg")

	// time.Sleep(5 * time.Second)

	// Loop:
	// 	for {
	// 		select {
	// 		case tempInt := <-data:
	// 			switch tempInt {
	// 			case 0:
	// 				break Loop
	// 			default:
	// 				fmt.Println(tempInt)
	// 			}
	// 		}
	// 	}
	fmt.Println(<-data)

	wg.Wait()

	fmt.Println("done wg 1 ")

	command := make(chan string)
	wg.Add(1)
	go routine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause" //inserting to channel string pause

	time.Sleep(1 * time.Second)
	command <- "Play" //inserting to channel string play

	time.Sleep(1 * time.Second)
	command <- "Stop" //inserting to channel string stop

	wg.Wait()
	fmt.Println("done wg 2 ")

	s = []string{"Python", "Java", "Golangggg"}
	for _, v := range s {
		wg.Add(1)
		// fmt.Printf("wg: %v\n", wg)
		go increment(v)

	}
	wg.Wait()
	fmt.Println("done wg 3 ")
	fmt.Println("Counter:", counter)

	//using mutex to single acces go routine
	counter = 0
	for _, v := range s {
		wg.Add(1)
		go increment2(v)

	}
	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)
	fmt.Println("done wg 4 ")
}

func responseSize(url string, data chan int) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: inserting")
	data <- len(body)
}

func routine(command chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	state := "Play"
	for { //infinite loop
		select {
		case cmd := <-command: //fill cmd with chan command string value
			fmt.Println("go routine state is : ", cmd)
			switch cmd {
			case "Pause":
				state = "Pause"
			case "Play":
				state = "Play"
			case "Stop":
				return //done go routine
			}
		default:
			if state == "Play" {
				doWork()
			}
		}

	}

}

var i int

func doWork() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func increment(name string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for range name {
		atomic.AddInt32(&counter, 1)
		//Race conditions occur due to unsynchronized access to shared resource and attempt to read and write to that resource at the same time.
		//Atomic functions provide low-level locking mechanisms for synchronizing access to integers and pointers.
		//Atomic functions generally used to fix the race condition.

		runtime.Gosched() // Yield the thread and be placed back in queue.
		//Gosched yields the processor, allowing other goroutines to run.
		//It does not suspend the current goroutine, so execution resumes automatically.
	}
}

func increment2(lang string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
		//A critical section defined by the calls to Lock() and Unlock()
		//protects the actions against the counter variable and reading the text of name variable.
	}
}
