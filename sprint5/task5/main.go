package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>
func isSearchTree(root *Node) ([]Node, bool) {

	if root == nil {
		return []Node{}, true
	}

	if root.left == nil && root.right == nil {
		return []Node{*root}, true
	}

	left, ok := isSearchTree(root.left)
	if !ok {
		return nil, false
	}

	right, ok := isSearchTree(root.right)
	if !ok {
		return nil, false
	}

	for _, node := range left {
		if node.value >= root.value {
			return nil, false
		}
	}

	for _, node := range right {
		if node.value <= root.value {
			return nil, false
		}
	}

	return append(left, right...), true
}

func Solution(root *Node) bool {
	_, ok := isSearchTree(root)
	return ok
}
