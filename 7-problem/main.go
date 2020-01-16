package main

import (
	"fmt"
	"strconv"
)

var duplicateNums = map[int]bool{
	11: true,
	12: true,
	13: true,
	14: true,
	15: true,
	16: true,
	17: true,
	18: true,
	19: true,
	21: true,
	22: true,
	23: true,
	24: true,
	25: true,
	26: true,
}

func main() {

	numbers := "12121213121212"
	fmt.Println(1 + totalPossibilities(numbers))
}

func totalPossibilities(numbers string) int {

	if len(numbers) <= 1 {
		return 0
	}

	num, err := strconv.Atoi(numbers[:2])
	if err != nil {
		return 0
	}

	if duplicateNums[num] {
		return 1 + totalPossibilities(numbers[2:]) + totalPossibilities(numbers[1:])
	}

	return 0 + totalPossibilities(numbers[1:])
}
