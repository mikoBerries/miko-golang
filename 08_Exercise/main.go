package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}
type person struct {
	First string
	Last  string
	Age   int
}
type UserVersion2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}
	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}
	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	//Hand on exercise #1
	//marshal the []user to JSON. There is a little bit of a curve ball in this
	//hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a
	//package.
	if myjson, err := json.Marshal(users); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(myjson))
	}
	//Hand on exercise #1
	//unmarshal the JSON into a Go data structure. Hint:
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	myUnmarshalJson := []person{}
	if err := json.Unmarshal([]byte(s), &myUnmarshalJson); err != nil {
		log.Fatal(err)
	}
	for _, v := range myUnmarshalJson {
		fmt.Println(v)
	}

	//Hands-on exercise #3
	//encode to JSON the []user sending the results to Stdout. Hint: you will
	//need to use json.NewEncoder(os.Stdout).encode(v interface{})

	user1 := UserVersion2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	user2 := UserVersion2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	user3 := UserVersion2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	tempSliceUser := []UserVersion2{user1, user2, user3}

	fmt.Println(users)

	// your code goes here
	ecd := json.NewEncoder(os.Stdout) //new json.encoder using os.stdout writer
	if err := ecd.Encode(tempSliceUser); err != nil {
		log.Fatal(err)
	}
	fmt.Println(tempSliceUser)

	/*
		Hands-on exercise #4
		Starting with this code, sort the []int and []string for each person.
	*/
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	/*
		Hands-on exercise #5
		Starting with this code, sort the []user by
			● age
			● last
		Also sort each []string “Sayings” for each user
			● print everything in a way that is pleasant

	*/

	myuser1 := UserVersion2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	myuser2 := UserVersion2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	myuser3 := UserVersion2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	dumpUser := []UserVersion2{myuser1, myuser2, myuser3}

	fmt.Println(users)

	// your code goes here

	sort.Sort(sortByFirst(dumpUser))

	for _, v := range dumpUser { //sorting saying per data
		fmt.Println(v)
		sort.Strings(v.Sayings) //sort saying first
		for _, v := range v.Sayings {
			fmt.Println(v)
		}
	}
}

type sortByFirst []UserVersion2

func (a sortByFirst) Len() int           { return len(a) }
func (a sortByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByFirst) Less(i, j int) bool { return a[i].First < a[j].First }
