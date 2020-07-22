package _04_arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("soma coleção de n... números", func(t *testing.T) {

		numbers := []int{1, 2, 3}
		want := 6
		got := Sum(numbers)

		if want != got {
			t.Errorf(" queria '%d', obteve '%d', dados %v", want, got, numbers)
		}
	})

	t.Run("soma n... coleções de n... números", func(t *testing.T) {
		want := []int{6, 9}
		got := SumAll([]int{1, 2, 3}, []int{0, 9})

		if !reflect.DeepEqual(want, got) {
			t.Errorf(" queria '%d', obteve '%d'", want, got)
		}
	})
}
