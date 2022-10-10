package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Person struct {
	First, Last string
	Age         int
}

func main() {
	jsonStream := `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`
	type Message struct {
		Name, Text string
	}
	myDecoder := json.NewDecoder(strings.NewReader(jsonStream))
	t, err := myDecoder.Token() // A Delim is a JSON array or object delimiter, one of [ ] { or }.
	if err != nil {             //check error
		log.Fatal(err)
	}
	fmt.Printf("%T \t %v", t, t)

	for myDecoder.More() { //has more value return bool T/F
		var m Message
		if err := myDecoder.Decode(&m); err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("%v: %v\n", m.Name, m.Text)

	}
	t, err = myDecoder.Token() // A Delim is a JSON array or object delimiter, one of [ ] { or }.
	if err != nil {            //check error
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n \n \n ", t, t)
	myPersons := []Person{{"agus", "tina", 24}, {"meta", "mimi", 25}}
	if result, err := json.Marshal(myPersons); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(result))
	}
}
