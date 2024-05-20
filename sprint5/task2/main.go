package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func isBalancedTree(root *Node) bool {
	if root == nil {
		return true
	}
	if !isBalancedTree(root.left) {
		return false
	}
	if level(root.left)-level(root.right) < -1 || level(root.left)-level(root.right) > 1 {
		return false
	}
	if !isBalancedTree(root.right) {
		return false
	}
	return true
}

func level(root *Node) int {
	if root == nil {
		return 0
	}
	return max(level(root.left), level(root.right)) + 1
}

func Solution(root *Node) bool {
	return isBalancedTree(root)
}

func test2() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{10, nil, nil}
	node5 := Node{2, &node3, &node4}
	if !Solution(&node5) {
		panic("WA")
	}

	//str := "0 0 1 2\n1 1 3 None\n2 2 None 4\n3 3 5 6\n4 4 7 8\n5 5 None None\n6 6 None None\n7 7 None None\n8 8 None None\n"
	//arr := strings.Fields(str)
	//for i := 0; i < len(arr); i++ {
	//	arrLine :=  strings.Fields(arr[2])
	//	for i := 0; i < len(arr); i++ {
	//		node := Node{value: 0}
	//	}
	//}

}
