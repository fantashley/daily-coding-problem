package main

import "fmt"

func main() {
	myString := []rune("my name is ashley")
	reverseString(myString)
	fmt.Println(string(myString))
}

func reverseString(r []rune) {
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
}
