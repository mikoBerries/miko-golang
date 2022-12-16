package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)

type customErr struct {
	function string
	line     string
	err      string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("there is error in %v line %v what happend %v", ce.function, ce.line, ce.err)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func main() {
	c1 := customErr{}
	c1.function = GetFunctionName(foo)
	c1.err = "this is what err happend"
	c1.line = "10000"
	foo(c1)
}

func foo(e error) {
	fmt.Printf("%v", e.(customErr).function) //type casting error to costumerr to get is secific val
	log.Printf("%v", e)
}
