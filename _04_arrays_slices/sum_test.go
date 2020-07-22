package _04_arrays_slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("soma coleção de n... números", func(t *testing.T) {

		numbers := []int{1, 2, 3}
		want := 6
		got := Sum(numbers)

		if want != got {
			t.Errorf(" queria '%d', obteve '%d', dados %v", want, got, numbers)
		}
	})
}
