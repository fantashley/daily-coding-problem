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
