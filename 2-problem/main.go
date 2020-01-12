package main

import "fmt"

func main() {

	numbers := []int{3, 2, 1}
	fmt.Println(productize(numbers))
}

func productize(numbers []int) []int {

	if len(numbers) <= 1 {
		return []int{0}
	}

	var products []int
	leftProducts := []int{numbers[0]}
	rightProducts := []int{numbers[len(numbers)-1]}

	for i := 1; i < len(numbers); i++ {
		leftProducts = append(leftProducts, numbers[i]*leftProducts[i-1])
	}

	for i := len(numbers) - 2; i >= 0; i-- {
		rightProducts = append([]int{numbers[i] * rightProducts[0]}, rightProducts...)
	}

	for i := range numbers {

		var leftProduct, rightProduct int

		if i == 0 {
			leftProduct = 1
		} else {
			leftProduct = leftProducts[i-1]
		}

		if i == len(numbers)-1 {
			rightProduct = 1
		} else {
			rightProduct = rightProducts[i+1]
		}

		products = append(products, leftProduct*rightProduct)
	}

	return products
}
