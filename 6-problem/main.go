/*
An XOR linked list is a more memory efficient doubly linked list. Instead of
each node holding next and prev fields, it holds a field named both, which is an
XOR of the next node and the previous node. Implement an XOR linked list; it has
an add(element) which adds the element to the end, and a get(index) which
returns the node at index.
*/

package main

import (
	"fmt"
	"unsafe"
)

type llnode struct {
	both *llnode
	data string
}

type linkedlist struct {
	beginning *llnode
}

func main() {

	myLinkedList := linkedlist{}
	myLinkedList.add("ashley")
	myLinkedList.add("is")
	myLinkedList.add("awesome")
	fmt.Println(myLinkedList.get(0))
	fmt.Println(myLinkedList.get(1))
	fmt.Println(myLinkedList.get(2))
}

func (ll *linkedlist) add(element string) {

	newNode := llnode{
		data: element,
	}

	if ll.beginning == nil {
		ll.beginning = &newNode
		return
	}

	prev := ll.beginning
	current := (*prev).both

	for current != nil {

		next := uintptr(unsafe.Pointer(prev)) ^ uintptr(unsafe.Pointer((*current).both))
		prev = current
		current = (*llnode)(unsafe.Pointer(next))
	}

	(*prev).both = (*llnode)(unsafe.Pointer(uintptr(unsafe.Pointer((*prev).both)) ^ uintptr(unsafe.Pointer(&newNode))))
	newNode.both = prev
}

func (ll *linkedlist) get(index int) llnode {

	prev := ll.beginning
	current := (*prev).both
	i := 0

	for (current != nil) && (i < index) {

		next := uintptr(unsafe.Pointer(prev)) ^ uintptr(unsafe.Pointer((*current).both))
		prev = current
		current = (*llnode)(unsafe.Pointer(next))
		i++
	}

	return *prev
}
