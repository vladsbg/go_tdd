package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSum := func(t *testing.T, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf(" queria '%d', obteve '%d'", want, got)
		}
	}

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

		checkSum(t, want, got)
	})

	t.Run("soma todo o resto", func(t *testing.T) {
		want := []int{2, 9}
		got := SumAllRest([]int{1, 2}, []int{0, 9})

		checkSum(t, want, got)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		want := []int{0, 9}
		got := SumAllRest([]int{}, []int{3, 4, 5})

		checkSum(t, want, got)
	})
}
