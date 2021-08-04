package main

import (
	"fmt"
	"sort"
)

// Given an array of time intervals (start, end) for classroom lectures
// (possibly overlapping), find the minimum number of rooms required.

// For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.

func main() {
	intervals := [][]int{
		{30, 75},
		{0, 50},
		{60, 150},
	}

	fmt.Printf("Minimum rooms needed is %d\n", minRooms(intervals))
}

func minRooms(intervals [][]int) int {
	var (
		startTimes = make([]int, len(intervals))
		endTimes   = make([]int, len(intervals))
	)

	for i, interval := range intervals {
		startTimes[i] = interval[0]
		endTimes[i] = interval[1]
	}

	sort.Ints(startTimes)
	sort.Ints(endTimes)

	var (
		startIndex  int
		endIndex    int
		overlaps    int
		maxOverlaps int
	)
	for startIndex < len(startTimes) && endIndex < len(endTimes) {
		if startTimes[startIndex] < endTimes[endIndex] {
			startIndex++
			overlaps++
			if overlaps > maxOverlaps {
				maxOverlaps = overlaps
			}
		} else if endTimes[endIndex] < startTimes[startIndex] {
			endIndex++
			overlaps--
		} else {
			startIndex++
			endIndex++
		}
	}

	return maxOverlaps
}
