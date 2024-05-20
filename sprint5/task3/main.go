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

type result []int

func (r *result) checkSum(root *Node, isLeft bool) {
	if root == nil {
		return
	}

	left := root.left
	right := root.right
	if isLeft {
		if left != nil {
			*r = append(*r, left.value)
		}

		if right != nil {
			*r = append(*r, right.value)
		}
	} else {

		if right != nil {
			*r = append(*r, right.value)
		}

		if left != nil {
			*r = append(*r, left.value)
		}
	}

	r.checkSum(left, true)
	r.checkSum(right, false)

	return
}

func Solution(root *Node) bool {

	rL := result{}
	rL.checkSum(root.left, true)

	rR := result{}
	rR.checkSum(root.right, false)

	if len(rL) != len(rR) {
		return false
	}

	for i := 0; i < len(rL); i++ {
		//if rL[i] != rR[len(rR)-1-i] {
		if rL[i] != rR[i] {
			return false
		}
	}

	return true
}

func main() {

	//node1 := Node{3, nil, nil}
	//node2 := Node{4, nil, nil}
	//node5 := Node{2, &node1, &node2}
	//
	//node3 := Node{4, nil, nil}
	//node4 := Node{3, nil, nil}
	//node6 := Node{2, &node3, &node4}
	//
	//node7 := Node{1, &node5, &node6}
	//
	//if !Solution(&node7) {
	//	print("error WA")
	//}

	input := "0 0 1 2\n1 2 3 None\n2 2 None 4\n3 3 5 6\n4 3 7 8\n5 0 None None\n6 1 None None\n7 1 None None\n8 0 None None"
	root := buildTreeFromString(input)
	print(Solution(root))

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
