package main

import (
	"fmt"

	"github.com/MikoBerries/miko-golang/13_exercise/02/quote"
	"github.com/MikoBerries/miko-golang/13_exercise/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
