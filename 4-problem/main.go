/*
Given an array of integers, find the first missing positive integer in linear
time and constant space. In other words, find the lowest positive integer that
does not exist in the array. The array can contain duplicates and negative
numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should
give 3.

You can modify the input array in-place.
*/

package main

import "fmt"

func main() {

	numbers := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPosInt(numbers))
}

func firstMissingPosInt(numbers []int) int {

	for i, n := range numbers {

		if n == 0 {
			numbers[i] = -1
		}
	}

	for i, n := range numbers {

		if i+1 == n {
			numbers[i] = 0
		} else if (n >= 1) && (n <= len(numbers)) {
			placeNumber(numbers, n)
		}
	}

	for i, n := range numbers {
		if n != 0 {
			return i + 1
		}
	}

	return len(numbers) + 1
}

func placeNumber(numbers []int, n int) {

	if (numbers[n-1] < 1) || (numbers[n-1] > len(numbers)) {
		numbers[n-1] = 0
	} else {
		placeNumber(numbers, numbers[n-1])
		numbers[n-1] = 0
	}
}
