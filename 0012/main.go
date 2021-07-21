package main

import "fmt"

/*
There exists a staircase with N steps, and you can climb up either 1 or 2 steps
at a time. Given N, write a function that returns the number of unique ways you
can climb the staircase. The order of the steps matters.

For example, if N is 4, then there are 5 unique ways:

    1, 1, 1, 1
    2, 1, 1
    1, 2, 1
    1, 1, 2
    2, 2

What if, instead of being able to climb 1 or 2 steps at a time, you could climb
any number from a set of positive integers X? For example, if X = {1, 3, 5}, you
could climb 1, 3, or 5 steps at a time.
*/

func main() {

	fmt.Println(uniqueClimbs(5, map[int]bool{
		1: true,
		3: true,
		5: true,
	}))
}

func uniqueClimbs(stairs int, steps map[int]bool) int {

	minStep := stairs

	var cache []int

	for step := range steps {
		if step < minStep {
			minStep = step
		}
	}

	for i := 0; i < minStep-1; i++ {
		cache = append(cache, 0)
	}

	cache = append(cache, 1)

	for i := minStep; i <= stairs; i++ {
		stepSum := 0
		for step := range steps {
			if i-step >= 0 {
				stepSum += cache[i-step]
			}
		}
		cache = append(cache, stepSum)
	}

	return cache[stairs]
}
