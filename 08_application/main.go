package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
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

	x := `[{"First":"agus","Last":"tina","Age":24},{"First":"meta","Last":"mimi","Age":25}]`
	byteSlice := []byte(x)

	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", byteSlice) //unt8 lias for byte
	fmt.Println("byte slice : ", byteSlice)

	tempPerson := []Person{}

	if err := json.Unmarshal(byteSlice, &tempPerson); err != nil {
		log.Fatal(err)
	}
	fmt.Println(tempPerson)

	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground") //os.Stdout is io writer
	io.WriteString(os.Stdout, "Hello, playground")

	//sorting

	sliceInt := []int{123, 941, 1, 134, 12345, 112, 3, 1231}
	sort.Ints(sliceInt)
	fmt.Println(sliceInt)

	sliceString := []string{"budi", "budak", "kuman", "kimak"}
	sort.Strings(sliceString)
	fmt.Println(sliceString)

	//costume sort in slice of struct

	p1 := Person{"James", "bondong", 32}
	p2 := Person{"Moneypenny", "money", 27}
	p3 := Person{"Q", "T", 64}
	p4 := Person{"M", "R", 56}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(sortByFirst(people))
	fmt.Println("first")
	fmt.Println(people)
	sort.Sort(sortByLast(people))
	fmt.Println("last")
	fmt.Println(people)
	sort.Sort(sortByAges(people))
	fmt.Println("ages")
	fmt.Println(people)

	//bycrpt password for masking password in DB
	stringPassword := "Test1234"
	tempbyte, err := bcrypt.GenerateFromPassword([]byte(stringPassword), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stringPassword)
	fmt.Println(tempbyte)
	inputPassword := "Test1234"
	if err := bcrypt.CompareHashAndPassword(tempbyte, []byte(inputPassword)); err != nil {
		log.Fatal(err) // hashedPassword is not the hash of the given password
	} else {
		fmt.Println("logged in")
	}

}

//sort.interface must implement 3 func and write your logic
//	Len() int
// Swap(i, j int)
// Less(i, j int) bool
type sortByFirst []Person

func (a sortByFirst) Len() int           { return len(a) }
func (a sortByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByFirst) Less(i, j int) bool { return a[i].First < a[j].First }

type sortByLast []Person

func (a sortByLast) Len() int           { return len(a) }
func (a sortByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

type sortByAges []Person

func (a sortByAges) Len() int           { return len(a) }
func (a sortByAges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByAges) Less(i, j int) bool { return a[i].Age < a[j].Age }
