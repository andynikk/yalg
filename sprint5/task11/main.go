package main

import "fmt"

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func printRange(root *Node, left int, right int) {
	if root == nil {
		return
	}

	printRange(root.left, left, right)

	if root.value >= left && root.value <= right {
		fmt.Println(root.value)
	}

	printRange(root.right, left, right)

}

func main() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, &node1}
	node3 := Node{8, nil, nil}
	node4 := Node{8, nil, &node3}
	node5 := Node{9, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node2, &node6}
	printRange(&node7, 2, 8)
	// expected output: 2 5 8 8
}
