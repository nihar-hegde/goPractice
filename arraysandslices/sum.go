package main

func Sum(numbers []int) int {
	sum := 0
	// NOTE: the range returns index and value below the _ indecates that the index is being ignored
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// NOTE  array is decalred as follows numbers[5] int
// NOTE slice is decalred as follows numbers[] int

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumber := len(numbersToSum)
	sums := make([]int, lengthOfNumber)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
