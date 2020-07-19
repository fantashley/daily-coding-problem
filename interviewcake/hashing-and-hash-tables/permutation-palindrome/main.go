package main

import "fmt"

func main() {
	fmt.Println(containsPalindrome("ivicc"))
}

func containsPalindrome(input string) bool {
	letterSet := make(map[rune]struct{})
	exists := struct{}{}
	for _, l := range input {
		if _, ok := letterSet[l]; ok {
			delete(letterSet, l)
		} else {
			letterSet[l] = exists
		}
	}
	return len(letterSet) <= 1
}
