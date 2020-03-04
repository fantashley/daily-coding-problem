package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Using a function rand5() that returns an integer from 1 to 5 (inclusive) with
uniform probability, implement a function rand7() that returns an integer from 1
to 7 (inclusive).
*/

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(rand7())
	}
}

func rand7() int {
	column := rand5()
	row := rand5()
	index := column + 5*(row-1)
	if index > 21 {
		return rand7()
	}
	return index%7 + 1
}

func rand5() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(5) + 1
}
