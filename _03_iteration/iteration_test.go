package _03_iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	want := "aaaaa"
	got := Repeat("a", 5)

	if want != got {
		t.Errorf(" queria '%s', obteve '%s'", want, got)
	}
}

func ExampleRepeat() {
	repeat := Repeat("c", 3)
	fmt.Println(repeat)
	// Output: ccc
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
