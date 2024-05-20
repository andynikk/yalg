package main

import (
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	u, v, w int
}

type DSU struct {
	parent []int
}

func newDSU(n int) *DSU {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{parent}
}

func (dsu *DSU) find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.find(dsu.parent[x])
	}
	return dsu.parent[x]
}

func (dsu *DSU) union(x, y int) {
	dsu.parent[dsu.find(x)] = dsu.find(y)
}

func minSpanningTree(n int, edges []Edge) string {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w > edges[j].w
	})

	dsu := newDSU(n)
	totalWeight := 0
	numEdges := 0

	for _, edge := range edges {
		if dsu.find(edge.u) != dsu.find(edge.v) {
			dsu.union(edge.u, edge.v)
			totalWeight += edge.w
			numEdges++
		}
	}

	result := "Oops! I did it again" + "\n"
	if numEdges == n-1 {
		result = strconv.Itoa(totalWeight) + "\n"
	}

	return result
}

func expensiveNetwork(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nm := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	lines = lines[1:]

	edges := make([]Edge, m)
	del := 0
	for i := 0; i < m; i++ {
		line := strings.Split(strings.Replace(lines[i], "\r", "", -1), " ")
		if line[0] == line[1] {
			del++
			continue
		}

		u, _ := strconv.Atoi(line[0])
		v, _ := strconv.Atoi(line[1])

		w, _ := strconv.Atoi(line[2])
		edges[i-del] = Edge{u - 1, v - 1, w}
	}

	edges = edges[:len(edges)-del]
	return minSpanningTree(n, edges)
}
