package main

import (
	"sort"
	"strconv"
	"strings"
)

func dfs(node int, graph map[int][]int, visited map[int]bool, component *[]string) {
	visited[node] = true
	*component = append(*component, strconv.Itoa(node))
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited, component)
		}
	}
}

func task05(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")

	verticesEdges := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")

	apexes, _ := strconv.Atoi(verticesEdges[0])
	edges, _ := strconv.Atoi(verticesEdges[1])

	input := lines[1 : edges+1]
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

	visited := make(map[int]bool)
	components := [][]string{}
	for node := 1; node <= apexes; node++ {
		if !visited[node] {
			component := []string{}
			dfs(node, graph, visited, &component)
			components = append(components, component)
		}
	}

	sort.Slice(components, func(i, j int) bool {
		i0, _ := strconv.Atoi(components[i][0])
		j0, _ := strconv.Atoi(components[j][0])
		return i0 < j0
	})

	result := []string{}
	for _, v := range components {
		sort.Slice(v, func(i, j int) bool {
			vI, _ := strconv.Atoi(v[i])
			vJ, _ := strconv.Atoi(v[j])
			return vI < vJ
		})
		result = append(result, strings.Join(v, " ")+" ")
	}

	return strconv.Itoa(len(components)) + "\n" + strings.Join(result, "\n") + "\n"
}
