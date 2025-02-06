package sum

func SumAll(numbersToSum ...[]int) []int {
	// lenOfNumbersToSum := len(numbersToSum)
	// sums := make([]int, lenOfNumbersToSum)
	var sums []int

	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		sums = append(sums, Sum(numbers))
	}

	return sums
}
