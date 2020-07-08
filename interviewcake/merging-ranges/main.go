package main

import (
	"fmt"
	"sort"
)

type meeting struct {
	startTime int
	endTime   int
}

func main() {
	meetingTimes := []meeting{
		{1, 10},
		{2, 6},
		{3, 5},
		{7, 9},
	}

	fmt.Println(mergeRanges(meetingTimes))
}

func mergeRanges(meetings []meeting) []meeting {
	var startTimes, endTimes []int
	for _, meeting := range meetings {
		startTimes = append(startTimes, meeting.startTime)
		endTimes = append(endTimes, meeting.endTime)
	}
	sort.Ints(startTimes)
	sort.Ints(endTimes)
	var mergedMeetings []meeting
	var startPos, endPos, currentStart, currentMeetings int
	for endPos < len(endTimes) {
		if (startPos == len(startTimes)) || (endTimes[endPos] < startTimes[startPos]) {
			currentMeetings--
			if currentMeetings == 0 {
				mergedMeetings = append(mergedMeetings, meeting{
					startTime: currentStart,
					endTime:   endTimes[endPos],
				})
			}
			endPos++
		} else if startTimes[startPos] < endTimes[endPos] {
			if currentMeetings == 0 {
				currentStart = startTimes[startPos]
			}
			currentMeetings++
			startPos++
		} else {
			startPos++
			endPos++
		}
	}
	return mergedMeetings
}
