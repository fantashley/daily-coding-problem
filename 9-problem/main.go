package main

import "fmt"

func main() {

	numbers := []int{5, 1, 1, 5}
	fmt.Println(largestSum(numbers))
}

func largestSum(numbers []int) int {

	if len(numbers) == 1 {
		return numbers[0]
	}

	nonadjacent := numbers[0]
	adjacent := largestSum(numbers[1:])

	if len(numbers) > 2 {
		nonadjacent += largestSum(numbers[2:])
	}

	if adjacent > nonadjacent {
		return adjacent
	}

	return nonadjacent
}
