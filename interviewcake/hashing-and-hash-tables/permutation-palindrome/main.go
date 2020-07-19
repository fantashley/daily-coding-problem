package main

import "fmt"

func main() {
	fmt.Println(containsPalindrome("civic"))
}

func containsPalindrome(input string) bool {
	letterMap := make(map[rune]int)
	var hasOdd bool
	for _, l := range input {
		letterMap[l]++
	}
	for _, count := range letterMap {
		if (count % 2) == 1 {
			if hasOdd {
				return false
			}
			hasOdd = true
		}
	}
	return true
}
