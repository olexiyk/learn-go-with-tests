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

func SumAllTails(setsOfNumbers ...[]int) []int {
	var sum []int
	for _, numbers := range setsOfNumbers {
		if len(numbers) > 0 {
			sum = append(sum, Sum(numbers[1:]))
		} else {
			sum = append(sum, 0)
		}
	}
	return sum
}

func Find(numbers []int, fn func(int) bool) (int, bool) {
	return 0, true
}

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
