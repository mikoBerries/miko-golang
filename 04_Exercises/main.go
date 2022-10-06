package main

import "fmt"

func main() {
	/*Hands-on exercise #1
	● Using a COMPOSITE LITERAL:
	● create an ARRAY which holds 5 VALUES of TYPE int
	● assign VALUES to each index position.
	● Range over the array and print the values out.
	● Using format printing
	○ print out the TYPE of the array
	*/

	myArr := make([]int, 5)
	for i, _ := range myArr {
		myArr[i] = i
	}
	for i, v := range myArr {
		fmt.Println(i, v)
	}

	/*
			Hands-on exercise #2
		● Using a COMPOSITE LITERAL:
		● create a SLICE of TYPE int
		● assign 10 VALUES
		● Range over the slice and print the values out.
		● Using format printing
		○ print out the TYPE of the slice
	*/
	myslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range myslice {
		fmt.Println(i, v)
	}

	/*	Hands-on exercise #3
		Using the code from the previous example, use SLICING to create the following new slices
		which are then printed:
		● [42 43 44 45 46]
		● [47 48 49 50 51]
		● [44 45 46 47 48]
		● [43 44 45 46 47]
	*/
	sliceOne := []int{42, 43, 44, 45, 46}
	sliceTwo := []int{47, 48, 49, 50, 51}
	sliceThree := []int{44, 45, 46, 47, 48}
	sliceFour := []int{43, 44, 45, 46, 47}

	slice2d := [][]int{sliceOne, sliceTwo, sliceThree, sliceFour}

	for _, v := range slice2d {
		fmt.Println(v)
	}

	/*
		Hands-on exercise #4
			Follow these steps:
			● start with this slice
			○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
			● append to that slice this value
			○ 52
			● print out the slice
			● in ONE STATEMENT append to that slice these values
			○ 53
			○ 54
			○ 55
			● print out the slice
			● append to the slice this slice
			○ y := []int{56, 57, 58, 59, 60}
			● print out the slice
	*/
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	/*Hands-on exercise #5
	To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise,
	follow these steps:
	● start with this slice
	○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	● use APPEND & SLICING to get these values here which you should ASSIGN to a
	variable “y” and then print:
	○ [42, 43, 44, 48, 49, 50, 51]
	*/
	x2 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y2 := append(x2[:3], x2[6:]...)
	fmt.Println(y2)

	/*
		Hands-on exercise #6
		Create a slice to store the names of all of the states in the United States of America. Use make
		and append to do this. Goal: do not have the array that underlies the slice created more than
		once. What is the length of your slice? What is the capacity? Print out all of the values, along
		with their index position, without using the range clause. Here is a list of the 50 states:
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
		Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
		Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
		Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
		Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,

	*/
	state := make([]string, 0, 50)
	fmt.Println(len(state))
	fmt.Println(cap(state))
	temp := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
	Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
	Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
	Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
	Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}
	// for i, v := range temp {
	// 	state[i] = v
	// }
	state = append(state, temp...)
	for i := 0; i < len(state); i++ {
		println(state[i])
	}

	fmt.Println(len(state))
	fmt.Println(cap(state))

	/*Hands-on exercise #7
		Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
	slice:
	"James", "Bond", "Shaken, not stirred"
	"Miss", "Moneypenny", "Helloooooo, James."
	Range over the records, then range over the data in each record.

	*/
	tempslice := []string{"James", "Bond", "Shaken, not stirred"}
	resultslice := [][]string{tempslice}
	tempslice = []string{"Miss", "Moneypenny", "Helloooooo, James."}
	resultslice = append(resultslice, tempslice)
	for _, v := range resultslice {
		fmt.Println(v)
	}
	/*
		Hands-on exercise #8
		Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
		TYPE []string which stores their favorite things. Store three records in your map. Print out all of
		the values, along with their index position in the slice.
		`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
	*/

	tempmap := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	// tempmap["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	// tempmap["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	// tempmap["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for k, v := range tempmap {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	for i, v := range tempmap {
		fmt.Println(i, v)
	}
	/*
		Hands-on exercise #9
		Using the code from the previous example, add a record to your map. Now print the map out
		using the “range” loop
	*/
	tempmap["new_things"] = []string{`belnded`, `corona`, `women?`}
	for i, v := range tempmap {
		fmt.Println(i, v)
	}

	/*
			Hands-on exercise #10
		Using the code from the previous example, delete a record from your map. Now print the map
		out using the “range” loop

	*/
	if _, ok := tempmap["new_things"]; ok {
		delete(tempmap, "new_things")
	}
	for i, v := range tempmap {
		fmt.Println(i, v)
	}
}
