package _04_arrays_slices

func Sum(numbers []int) (result int) {
	for _, value := range numbers {
		result += value
	}
	return
}
