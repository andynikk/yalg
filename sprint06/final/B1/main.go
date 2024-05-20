package main

import (
	"strconv"
	"strings"
)

const (
	B = iota
	R
)

func typeRoad(roadType uint8) int {
	if roadType == 'B' {
		return B
	}

	return R
}

func dfs(startNode int, roads [][]int, colors []int) bool {
	stack := []int{}
	stack = append(stack, startNode)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if colors[node] == 0 {
			colors[node] = 1
			stack = append(stack, node)

			for i := 0; i < len(roads[node]); i++ {
				if colors[roads[node][i]] == 0 {
					stack = append(stack, roads[node][i])
				} else if colors[roads[node][i]] == 1 {
					return true
				}
			}
		} else if colors[node] == 1 {
			colors[node] = 2
		}
	}
	return false
}

func checkRoads(cities int, roads [][]int) bool {
	colors := make([]int, cities+1)
	for node := 1; node < len(colors); node++ {
		if colors[node] != 0 {
			continue
		}
		if dfs(node, roads, colors) {
			return false
		}
	}
	return true
}

func railways(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	n, _ := strconv.Atoi(lines[0])
	lines = lines[1:]

	roads := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		cityOut := i + 1
		for j := 0; j < len(lines[i]); j++ {
			cityIn := cityOut + j + 1
			if typeRoad(lines[i][j]) == B {
				roads[cityOut] = append(roads[cityOut], cityIn)
			} else {
				roads[cityIn] = append(roads[cityIn], cityOut)
			}
		}
	}

	if checkRoads(n, roads) {
		return "YES\n"
	} else {
		return "NO\n"
	}
}
