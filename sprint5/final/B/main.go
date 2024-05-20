package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func remove(node *Node, key int) *Node {
	// Your code
	// “ヽ(´▽｀)ノ”

	if node == nil {
		return nil
	}

	if key < node.value {
		node.left = remove(node.left, key)
	} else if key > node.value {
		node.right = remove(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		minRight := findMin(node.right)
		node.value = minRight.value
		node.right = remove(node.right, minRight.value)
	}

	return node

}

func findMin(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func finalB() {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}
	newHead := remove(&node7, 10)
	if newHead.value != 5 {
		panic("WA")
	}
	if newHead.right != &node5 {
		panic("WA")
	}
	if newHead.right.value != 8 {
		panic("WA")
	}
}
