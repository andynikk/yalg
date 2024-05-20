package main

import "math"

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func maxi(root *Node, m int) int {
	if root == nil {
		return m
	}

	m = maxi(root.left, max(m, root.value))
	m = maxi(root.right, max(m, root.value))

	return m
}

func Solution(root *Node) int {
	// Your code
	// “ヽ(´▽｀)ノ”

	if root == nil {
		return math.MinInt
	}

	// <template>
	return maxi(root, math.MinInt)
}

func test1() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, &node3, nil}
	if Solution(&node4) != 3 {
		panic("WA")
	}
}
