package _04_arrays_slices

func Sum(numbers []int) (result int) {
	for _, value := range numbers {
		result += value
	}
	return
}

func SumAll(numbers ...[]int) (result []int) {
	countNumbers := len(numbers)
	result = make([]int, countNumbers)

	for i, values := range numbers {
		result[i] = Sum(values)
	}
	return
}
