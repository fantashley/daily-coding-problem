package main

import "fmt"

func main() {
	myList := []int{3, 4, 6, 10, 11, 15}
	alicesList := []int{1, 5, 8, 12, 14, 19}
	fmt.Println(mergeArrays(myList, alicesList))
}

func mergeArrays(array1, array2 []int) []int {
	var mergedArray []int
	var index1, index2 int
	for {
		if index1 < len(array1) && (array1[index1] < array2[index2]) {
			mergedArray = append(mergedArray, array1[index1])
			index1++
		} else if index2 < len(array2) {
			mergedArray = append(mergedArray, array2[index2])
			index2++
		} else {
			break
		}
	}
	return mergedArray
}
