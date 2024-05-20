package main

import (
	"sort"
	"strconv"
	"strings"
)

const (
	white = iota
	grey
	black
)

func dfs(arr [][]int, start int, apexes int) string {
	color := make([]int, apexes+1)
	stack := []int{start}

	result := []string{}

	for len(stack) > 0 {

		apex := stack[len(stack)-1]
		stack = (stack)[:len(stack)-1]

		if color[apex] == white {
			result = append(result, strconv.Itoa(apex))
			color[apex] = 1
			stack = append(stack, apex)
			for _, adjacent := range arr[apex] {
				if color[adjacent] == 0 {
					stack = append(stack, adjacent)
				}
			}
		} else if color[apex] == grey {
			color[apex] = black
		}
	}
	return strings.Join(result, " ") + " "
}

func task03(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")

	verticesEdges := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")

	apexes, _ := strconv.Atoi(verticesEdges[0])
	edges, _ := strconv.Atoi(verticesEdges[1])

	input := lines[1 : edges+1]
	arrInput := make([][]int, apexes+1)
	for _, v := range input {
		if v == "" {
			continue
		}

		tmp := strings.Split(strings.Replace(v, "\r", "", -1), " ")
		edgeOut, _ := strconv.Atoi(tmp[0])
		edgeIn, _ := strconv.Atoi(tmp[1])

		arrInput[edgeOut] = append(arrInput[edgeOut], edgeIn)
		arrInput[edgeIn] = append(arrInput[edgeIn], edgeOut)
	}

	for i := 0; i < len(arrInput); i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(arrInput[i])))
	}

	startApex, _ := strconv.Atoi(lines[1+len(input)])
	result := dfs(arrInput, startApex, apexes)

	return result
}
