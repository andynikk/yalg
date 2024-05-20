package main

import (
	"sort"
	"strconv"
	"strings"
)

func bfs(graph map[int][]int, start int) []string {
	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true
	result := []string{strconv.Itoa(start)}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				result = append(result, strconv.Itoa(neighbor))
			}
		}
	}

	return result
}

func task04(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	//verticesEdges := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")

	//apexes, _ := strconv.Atoi(verticesEdges[0])
	//edges, _ := strconv.Atoi(verticesEdges[1])

	input := lines[1 : len(lines)-1]
	graph := make(map[int][]int)

	for _, v := range input {
		if v == "" {
			continue
		}

		tmp := strings.Split(strings.Replace(v, "\r", "", -1), " ")
		edgeOut, _ := strconv.Atoi(tmp[0])
		edgeIn, _ := strconv.Atoi(tmp[1])

		graph[edgeOut] = append(graph[edgeOut], edgeIn)
		graph[edgeIn] = append(graph[edgeIn], edgeOut)
	}

	for _, v := range graph {
		sort.Sort(sort.IntSlice(v))
	}

	start, _ := strconv.Atoi(lines[len(lines)-1])
	result := bfs(graph, start)

	return strings.Join(result, " ") + " "
}
