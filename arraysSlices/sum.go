package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(intSlices ...[]int) (sums []int) {
	for _, ints := range intSlices {
		sums = append(sums, Sum(ints))
	}
	return
}

func SumAllTails(intSlices ...[]int) (sums []int) {
	for _, ints := range intSlices {
		if len(ints) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(ints[1:]))
		}
	}
	return
}
