package main

import (
	"strconv"
	"strings"
)

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func checkNodeRight(node1 *Node, node2 *Node) bool {

	if node1 == nil && node2 == nil {
		return true
	}

	if (node1 != nil && node2 == nil) || (node1 == nil && node2 != nil) {
		return false
	}

	if node1.value != node2.value {
		return false
	}

	return checkNodeRight(node1.right, node2.right)
}

func checkNodeLeft(node1 *Node, node2 *Node) bool {

	if node1 == nil && node2 == nil {
		return true
	}

	if (node1 != nil && node2 == nil) || (node1 == nil && node2 != nil) {
		return false
	}

	if node1.value != node2.value {
		return false
	}

	return checkNodeLeft(node1.left, node2.left)
}

func Solution(root1 *Node, root2 *Node) bool {
	// Your code
	// “ヽ(´▽｀)ノ”

	if root1.left == nil && root2.left == nil && root1.right == nil && root2.right == nil {
		return root1.value == root2.value
	}

	return checkNodeLeft(root1.left, root2.left) && checkNodeRight(root1.right, root2.right)
}

func main() {
	//node1 := Node{1, nil, nil}
	//node2 := Node{2, nil, nil}
	//node3 := Node{3, &node1, &node2}
	//node4 := Node{1, nil, nil}
	//node5 := Node{2, nil, nil}
	//node6 := Node{3, &node4, &node5}
	//
	//if !Solution(&node3, &node6) {
	//	println("WA")
	//}

	input1 := "0 1 None None"
	root1 := buildTreeFromString(input1)

	input2 := "0 0 None None"
	root2 := buildTreeFromString(input2)

	print(Solution(root1, root2))

	print("\n")
	print("-----")
	print("\n")

	input1 = "0 0 1 2\n1 2 3 None\n2 2 None 4\n3 3 5 6\n4 3 7 8\n5 0 None None\n6 1 None None\n7 1 None None\n8 0 None 9\n9 1 None None"
	root1 = buildTreeFromString(input1)

	input2 = "0 0 1 2\n1 2 3 None\n2 2 None 4\n3 3 5 6\n4 3 7 8\n5 0 None None\n6 1 None None\n7 1 None None\n8 0 None None"
	root2 = buildTreeFromString(input2)

	print(Solution(root1, root2))

}

func buildTreeFromString(input string) *Node {

	lines := strings.Split(input, "\n")
	arrNod := make([][3]int, len(lines))
	nodes := make([]Node, len(lines))
	for i := 0; i < len(lines); i++ {

		fields := strings.Fields(lines[i])
		if len(fields) < 4 {
			continue
		}

		v, _ := strconv.Atoi(fields[1])
		l := -1
		if fields[2] != "None" {
			l, _ = strconv.Atoi(fields[2])
		}

		r := -1
		if fields[3] != "None" {
			r, _ = strconv.Atoi(fields[3])
		}

		arrNod[i] = [3]int{v, l, r}
		nodes[i] = Node{value: v, left: nil, right: nil}
	}

	for i := 0; i < len(nodes); i++ {

		if arrNod[i][1] == -1 {
			nodes[i].left = nil
		} else {
			nodes[i].left = &nodes[arrNod[i][1]]
		}

		if arrNod[i][2] == -1 {
			nodes[i].right = nil
		} else {
			nodes[i].right = &nodes[arrNod[i][2]]
		}
	}

	return &nodes[0]
}
