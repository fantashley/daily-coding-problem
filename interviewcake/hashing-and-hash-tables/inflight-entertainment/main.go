package main

import "fmt"

func main() {
	movieLengths := []int{10, 20, 30, 40, 45}
	fmt.Println(getMovies(90, movieLengths))
}

func getMovies(flightLength int, movieLengths []int) bool {
	seenLengths := make(map[int]struct{})
	var exists = struct{}{}
	for _, l := range movieLengths {
		if l >= flightLength {
			continue
		} else if _, ok := seenLengths[flightLength-l]; ok {
			return true
		}
		seenLengths[l] = exists
	}
	return false
}
