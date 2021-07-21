package main

import "fmt"

func main() {

	numbers := []int{5, 1, 1, 5}
	fmt.Println(largestSum(numbers))
}

func largestSum(numbers []int) int {

	switch l := len(numbers); l {
	case 0:
		return 0
	case 1:
		return numbers[0]
	case 2:
		if numbers[0] > numbers[1] {
			return numbers[0]
		}
		return numbers[1]
	}

	var adjMax int
	nonadjMax := numbers[0]

	if numbers[1] > numbers[0] {
		adjMax = numbers[1]
	} else {
		adjMax = numbers[0]
	}

	for i := 2; i < len(numbers); i++ {
		currentMax := nonadjMax + numbers[i]
		if currentMax > adjMax {
			nonadjMax = adjMax
			adjMax = currentMax
		} else {
			nonadjMax = adjMax
		}
	}

	return adjMax
}
