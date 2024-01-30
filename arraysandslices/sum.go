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
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
		//NOTE: the slices can be sliced numbers[1:] will return every value form 1 to end [from:to] or [low:heigh]
	}
	return sums
}
