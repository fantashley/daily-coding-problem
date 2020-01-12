package main

func productize(numbers []int) []int {
	if len(numbers) <= 1 {
		return []int{0}
	}
	var products []int
	leftProducts := []int{numbers[0]}
	rightProducts := []int{numbers[len(numbers)-1]}
	for i := 1; i < len(numbers); i++ {
		leftProducts = append(leftProducts, numbers[i]*leftProducts[i-1])
		rightProducts = append(numbers[len(numbers)-1-i]*rightProducts[len(numbers)-i], rightProducts)
	}
	for i, n := range numbers {

	}
}
