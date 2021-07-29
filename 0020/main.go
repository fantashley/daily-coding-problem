package main

import "log"

// Given two singly linked lists that intersect at some point, find the
// intersecting node. The lists are non-cyclical.

// For example, given A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10, return
// the node with value 8.

// In this example, assume nodes with the same value are the exact same node
// objects.

// Do this in O(M + N) time (where M and N are the lengths of the lists) and
// constant space.

func main() {
	aNodes := []int{5, 9, 3, 7, 10, 85}
	bNodes := []int{99, 1, 10, 85}
	nodeMap := make(map[int]*Node)

	for _, nodeNum := range append(aNodes, bNodes...) {
		nodeMap[nodeNum] = &Node{Value: nodeNum}
	}

	linkNodes(aNodes, nodeMap)
	linkNodes(bNodes, nodeMap)

	node := findIntersection(nodeMap[3], nodeMap[99])

	if node == nil {
		log.Fatal("Failed to find intersection node")
	}

	log.Printf("Intersecting node is %d", node.Value)
}

type Node struct {
	Value int
	Next  *Node
}

func findIntersection(A, B *Node) *Node {
	lenA := length(A)
	lenB := length(B)

	shortList := A
	longList := B
	if lenA > lenB {
		shortList = B
		longList = A
	}

	for i := 0; i < abs(lenA-lenB); i++ {
		longList = longList.Next
	}

	for shortList != longList {
		shortList = shortList.Next
		longList = longList.Next
	}

	return shortList
}

func linkNodes(nodeList []int, nodeMap map[int]*Node) {
	for i := 0; i < len(nodeList)-1; i++ {
		currNode := nodeList[i]
		if nodeMap[currNode].Next != nil {
			return
		}

		nextNode := nodeList[i+1]
		nodeMap[currNode].Next = nodeMap[nextNode]
	}
}

func length(list *Node) int {
	var total int
	currentNode := list

	for currentNode != nil {
		total++
		currentNode = currentNode.Next
	}

	return total
}

func abs(num int) int {
	if num >= 0 {
		return num
	}

	return -1 * num
}
