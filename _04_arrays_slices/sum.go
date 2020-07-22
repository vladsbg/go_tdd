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

func SumAllRest(numbers ...[]int) (result []int) {
	for _, values := range numbers {
		if len(values) > 0 {
			final := values[1:]
			result = append(result, Sum(final))
		} else {
			result = append(result, 0)
		}
	}

	return
}
