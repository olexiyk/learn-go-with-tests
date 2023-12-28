package main

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(setsOfNumbers ...[]int) []int {
	var sum []int
	for _, numbers := range setsOfNumbers {
		sum = append(sum, Sum(numbers))
	}
	return sum
}
