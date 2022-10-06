package main

import (
	"fmt"
)

type person struct {
	firstName, lastName      string
	favoriteIceCreamFlavours []string
}

type vehicle struct {
	door  int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	/*
			Hands-on exercise #1
		Create your own type “person” which will have an underlying type of “struct” so that it can store
		the following data:
		● first name
		● last name
		● favorite ice cream flavors
		Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
		which stores the favorite flavors.
	*/
	person1 := person{
		"agus", "budiman", []string{"vanila", "mocha"},
	}
	person2 := person{
		"budi", "handuk", []string{"coklat", "pisang"},
	}
	fmt.Println(person1)
	fmt.Println(person1.favoriteIceCreamFlavours)
	fmt.Println(person2)
	fmt.Println(person2.favoriteIceCreamFlavours)
	/*
		Hands-on exercise #2
		Take the code from the previous exercise, then store the values of type person in a map with the
		key of last name. Access each value in the map. Print out the values, ranging over the slice.

	*/
	m := map[string]person{}
	m[person1.lastName] = person1
	m[person2.lastName] = person2
	for _, v := range m {
		fmt.Println(v.firstName, v.lastName)
		for _, v := range v.favoriteIceCreamFlavours {
			fmt.Printf("\t" + v + "\n")
		}
	}

	/*
			Hands-on exercise #3
		● Create a new type: vehicle.
		○ The underlying type is a struct.
		○ The fields:
		■ doors
		■ color
		● Create two new types: truck & sedan.
		○ The underlying type of each of these new types is a struct.
		○ Embed the “vehicle” type in both truck & sedan.
		○ Give truck the field “fourWheel” which will be set to bool.
		○ Give sedan the field “luxury” which will be set to bool. solution
		● Using the vehicle, truck, and sedan structs:
		○ using a composite literal, create a value of type truck and assign values to the
		fields;
		○ using a composite literal, create a value of type sedan and assign values to the
		fields.
		● Print out each of these values.
		● Print out a single field from each of these values.

	*/
	mytruck := truck{
		vehicle{4, "black"},
		true,
	}
	//or
	mySedan := sedan{
		vehicle: vehicle{2, "white"},
		luxury:  true,
	}
	fmt.Println(mytruck.door, mytruck.color, mytruck.fourWheel)
	fmt.Println(mySedan.door, mySedan.color, mySedan.luxury)

	/*
		Hands-on exercise #4
		Create and use an anonymous struct
	*/

	something := struct {
		name                   string
		value1, value2, value3 int
		isOK                   bool
	}{
		"agus", 1, 2, 3, true,
	}
	fmt.Println(something)
}
