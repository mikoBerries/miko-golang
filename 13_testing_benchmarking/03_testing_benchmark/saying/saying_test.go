//Testing saying file and func
package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Jamesbond")
	if s != "Welcome my dear Jamesbond" {
		t.Error("got", s, "expected", "Welcome my dear Jamesbond")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Jamesbond"))
	//Output:
	//Welcome my dear Jamesbond

}

func BenchmarkGreet(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Greet("Jamesbond")
	}
}
