package main

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func main() {

	root := treeNode{
		val: 0,
		left: &treeNode{
			val: 1,
		},
		right: &treeNode{
			val: 0,
			left: &treeNode{
				val: 1,
				left: &treeNode{
					val: 1,
				},
				right: &treeNode{
					val: 1,
				},
			},
			right: &treeNode{
				val: 0,
			},
		},
	}

	_, totalUnivals := countUnivals(root)

	fmt.Println(totalUnivals)
}

func countUnivals(treeRoot treeNode) (bool, int) {

	if (treeRoot.left == nil) && (treeRoot.right == nil) {
		return true, 1
	}

	leftUnival, leftTotal := countUnivals(*(treeRoot.left))
	rightUnival, rightTotal := countUnivals(*(treeRoot.right))

	if leftUnival && rightUnival {

		if (treeRoot.left.val == treeRoot.right.val) && (treeRoot.left.val == treeRoot.val) {
			return true, leftTotal + rightTotal + 1
		}
	}

	return false, leftTotal + rightTotal
}
