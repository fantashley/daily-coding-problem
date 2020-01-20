package main

func main() {

	numbers := []int{2, 4, 6, 2, 5}

}

func largestSum(numbers []int) int {

	adjacent := largestSum(numbers[1:])
	nonadjacent := largestSum(numbers[2:])
}
