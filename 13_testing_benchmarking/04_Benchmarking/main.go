package main

import (
	"fmt"
	"strings"
)

func main() {
	dummyText := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	splitted := strings.Split(dummyText, " ")
	fmt.Println(EfficientCat(splitted, " "))
	fmt.Println(CompetingCat(splitted, " "))
}

func CompetingCat(xs []string, sep string) string {

	if len(xs) <= 0 {
		return ""
	}

	sizeOfArrayNeeded := len(sep) * (len(xs) - 1)
	for i := 0; i < len(xs); i++ {
		sizeOfArrayNeeded += len(xs[i])
	}

	b := make([]byte, sizeOfArrayNeeded)
	bp := copy(b, xs[0])

	for _, x := range xs[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], x)
	}
	return string(b)

}

func Joiner(xs []string, sep string) string {
	return strings.Join(xs, sep)
}

func EfficientCat(xs []string, sep string) string {
	switch len(xs) {
	case 0:
		return ""
	case 1:
		return xs[0]
	case 2:
		return xs[0] + sep + xs[1]
	case 3:
		return xs[0] + sep + xs[1] + sep + xs[2]
	}
	sizeOfArrayNeeded := len(sep) * (len(xs) - 1)
	for i := 0; i < len(xs); i++ {
		sizeOfArrayNeeded += len(xs[i])
	}

	b := make([]byte, sizeOfArrayNeeded)
	bp := copy(b, xs[0])

	for _, x := range xs[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], x)
	}
	return string(b)

}
