package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	s := Years(10)
	if s != 70 {
		t.Error("Get", s, "Expected", 70)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	//output:
	//70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(100)
	}
}

func TestYearstwo(t *testing.T) {
	s := YearsTwo(10)
	if s != 70 {
		t.Error("Get", s, "Expected", 70)
	}
}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//output:
	//70
}

func BenchmarkYearstwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(100)
	}
}
