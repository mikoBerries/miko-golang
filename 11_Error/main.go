package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}
func main() {

	logFile, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile) // log string will be output in spesific file path
	_, err = os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err) //standart error print
		log.Println("err happened", err) //Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
		//		log.Fatalln(err)//the Fatal functions call os.Exit(1) after writing the log message ...
		//		log.Panicln(err)//Panicln is equivalent to Println() followed by a call to panic().
		//		panic(err)
	}
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))

	//costuming error string
	f2()
	fmt.Println("Returned normally from f.")

	// e := errors.New()
	_, err = sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("norgate math: square root of negative number") //givin new error with own string

		nme := fmt.Errorf("norgate math again: square root of negative number: %v", f) //fmt,error fromat
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}                      //norgateMathError class clompeting error() string interface
	}
	return 42, nil
}
func f2() {
	defer func() {
		if r := recover(); r != nil { // recover() get value from panic
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i)) // panic exiting func but defer still executed
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
