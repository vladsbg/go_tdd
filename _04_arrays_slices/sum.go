package _04_arrays_slices

func Sum(numbers []int) (result int) {
	for _, value := range numbers {
		result += value
	}
	return
}

func SumAll(numbers ...[]int) (result []int) {
	for _, values := range numbers {
		result = append(result, Sum(values))
	}
	return
}
