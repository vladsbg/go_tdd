package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	result := Addition(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("esperado '%d', resultado '%d'", expected, result)
	}
}

func ExampleAddition() {
	soma := Addition(1, 3)
	fmt.Println(soma)
	// Output: 6
}
