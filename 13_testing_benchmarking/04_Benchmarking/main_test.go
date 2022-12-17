package main

import (
	"testing"
)

var x5 []string = []string{"Although", "this", "should", "be", "faster"}
var x3 []string = []string{"Although", "this", "should"}
var x2 []string = []string{"Although", "this"}
var x1 []string = []string{"Although"}
var x0 []string = []string{}

func TestEfficientCat(t *testing.T) {

}

func TestCompetingCat(t *testing.T) {

}

func ExampleEfficientCat() {}

func ExampleCompetingCat() {}

// For 5 Strings

func BenchmarkEfficientCat5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientCat(x5, " ")
	}
}

func BenchmarkCompetingCat5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompetingCat(x5, " ")
	}
}

func BenchmarkJoiner5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Joiner(x5, " ")
	}
}

// For 3 String

func BenchmarkEfficientCat3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientCat(x3, " ")
	}
}

func BenchmarkCompetingCat3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompetingCat(x3, " ")
	}
}

func BenchmarkJoiner3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Joiner(x3, " ")
	}
}

//For 2 Strings

func BenchmarkEfficientCat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientCat(x2, " ")
	}
}

func BenchmarkCompetingCat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompetingCat(x2, " ")
	}
}

func BenchmarkJoiner2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Joiner(x2, " ")
	}
}

// For 1 String

func BenchmarkEfficientCat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientCat(x1, " ")
	}
}

func BenchmarkCompetingCat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompetingCat(x1, " ")
	}
}

func BenchmarkJoiner1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Joiner(x1, " ")
	}
}

//For 0 Stirngs

func BenchmarkEfficientCat0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientCat(x0, " ")
	}
}

func BenchmarkCompetingCat0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompetingCat(x0, " ")
	}
}

func BenchmarkJoiner0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Joiner(x0, " ")
	}
}
