package main

import "fmt"

func main() {
	//array
	var x [5]int
	fmt.Println(len(x))
	x[4] = 0

	//slice
	//x := type{values} //composite literal
	y := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	//acces slice

	for i, v := range y {
		fmt.Println(i, v)
	}
	// colon in slice index

	//{1, 2, 3, 4}
	// 0  1  2  3
	fmt.Println(y[1:])  //start from index 1 till end
	fmt.Println(y[1:4]) //start from index 1 until <4

	//append slice and slice
	// .... unlimite ddata

	y = append(y, 11, 12, 13, 14, 15)
	fmt.Println(y)
	z := []int{123, 456, 789}
	y = append(y, z...)
	fmt.Println(y)

	//deleting data from slice
	y = append(y[:2], y[4:]...)
	fmt.Println(y)

	//make (type,len, cap)
	//array is slice with cap
	xyz := make([]int, 10, 20)
	xyz = append(xyz, y[:11]...)
	fmt.Println(xyz)
	fmt.Println(len(xyz))
	fmt.Println(cap(xyz))

	//multidimentional array

	onePeople := []string{"miko", "berries", "087"}
	twoPeople := []string{"miko2", "berries2", "0872222"}
	crowdofpeople := [][]string{onePeople, twoPeople}
	fmt.Println(crowdofpeople)

	//map
	m := map[string]int{
		"jamesbond": 100,
		"bondjames": 10,
	}
	fmt.Println(m["jamesbond"])
	fmt.Println(m["bondjames"])
	fmt.Println(m["jamesbondong"]) // unfind index return default int "0"

	if v, ok := m["jamesbondong"]; ok {
		fmt.Println(v)
		fmt.Println(ok)
	} else {
		fmt.Println(v)
		fmt.Println(ok)
	}
	//adding elemnt to map

	m["newThings"] = 32000 //adding new idex adn new value
	m["deletethis"] = 9999999
	for i, v := range m {
		println(i, v)
	}
	//deleting map

	delete(m, "deletethis")
	for i, v := range m {
		println(i, v)
	}

}