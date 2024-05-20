package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	res := "0 5 1 2\n1 3 3 4\n2 8 5 6\n3 1 None None\n4 3 None None\n5 6 None None\n6 9 None None"
	out := false
	answer := Solution(buildTreeFromString(res))
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = "0 5 1 2\n1 3 3 4\n2 8 5 6\n3 1 None None\n4 4 None None\n5 6 None None\n6 9 None None"
	out = true
	answer = Solution(buildTreeFromString(res))
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = "0 5 1 2\n1 4 3 4\n2 6 5 6\n3 2 None None\n4 6 None None\n5 4 None None\n6 8 None None"
	out = false
	answer = Solution(buildTreeFromString(res))
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = "0 5 1 2\n1 3 3 4\n2 8 5 6\n3 1 None None\n4 4 None None\n5 6 None None\n6 9 None 7\n7 19 None None"
	out = false
	answer = Solution(buildTreeFromString(res))
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}
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
