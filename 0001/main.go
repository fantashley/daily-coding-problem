/*
Given a list of numbers and a number k, return whether any two numbers from the
list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{10, 15, 3, 7}
	k := 17

	fmt.Println(addsToK(numbers, k))
}

func addsToK(numbers []int, k int) bool {
	sort.Ints(numbers)
	for _, n := range numbers {
		if binarySearch(numbers, 0, len(numbers)-1, k-n) != -1 {
			return true
		}
	}
	return false
}

func binarySearch(numbers []int, l int, r int, x int) int {

	for l <= r {
		mid := (l + r) / 2

		if x < numbers[mid] {
			return binarySearch(numbers, l, mid-1, x)
		} else if x > numbers[mid] {
			return binarySearch(numbers, mid+1, r, x)
		} else {
			return numbers[mid]
		}
	}

	return -1
}
