package word

import (
	"fmt"
	"testing"

	"github.com/MikoBerries/miko-golang/13_exercise/02/quote"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two two three")
	for key, value := range m {
		switch key {
		case "one":
			if value != 1 {
				t.Error("GOT", value, "Expected", 1)
			}
		case "two":
			if value != 2 {
				t.Error("GOT", value, "Expected", 2)
			}
		case "three":
			if value != 1 {
				t.Error("GOT", value, "Expected", 1)
			}
		}
	}
}
func TestCount(t *testing.T) {
	n := Count("ABC DEF XYZ")
	if n != 3 {
		t.Error("GOT", n, "Expected", 3)
	}
}

func BenchmarkUsecount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleUseCount() {
	fmt.Println(Count("ABC DEF XYZ"))
	//output:
	//3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
