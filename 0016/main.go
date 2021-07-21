/*
You run an e-commerce website and want to record the last N order ids in a log.
Implement a data structure to accomplish this, with the following API:

    record(order_id): adds the order_id to the log
    get_last(i): gets the ith last element from the log. i is guaranteed to be smaller than or equal to N.

You should be as efficient with time and space as possible.
*/
package main

import "fmt"

type log struct {
	orderIDs []int
	index    int
}

func main() {
	N := 10

	l := log{
		orderIDs: make([]int, N),
		index:    0,
	}

	for i := 0; i < 15; i++ {
		l.record(i)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(l.getLast(i))
	}
}

func (l *log) record(orderID int) {
	l.orderIDs[l.index] = orderID
	l.index = (l.index + 1) % len(l.orderIDs)
}

func (l *log) getLast(i int) int {
	if i <= l.index {
		return l.orderIDs[l.index-i]
	}
	n := len(l.orderIDs)
	return l.orderIDs[n-(i-l.index)]
}
