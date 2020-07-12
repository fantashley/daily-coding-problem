package main

import "fmt"

func main() {
    message := []rune("cake pound steal")
    reverseWords(message)
    fmt.Println(string(message))
}

func reverseWords(message []rune) {
	reverseString(message)
    var firstIndex, currentIndex int
    for firstIndex <= len(message) {
        if currentIndex == len(message) || message[currentIndex] == ' ' {
            reverseString(message[firstIndex:currentIndex])
			firstIndex = currentIndex + 1
		}
        currentIndex++
    }
}

func reverseString(word []rune) {
    for i := 0; i < len(word)/2; i++ {
        tmp := word[i]
        word[i] = word[len(word)-1-i]
        word[len(word)-1-i] = tmp
    }
}
