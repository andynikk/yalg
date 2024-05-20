package main

import (
	"strconv"
	"strings"
)

func makeArrString(c int) []string {
	a := []string{}
	for i := 0; i < c; i++ {
		a = append(a, "0")
	}

	return a
}

func task02(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")

	verticesEdges := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")
	vertices, _ := strconv.Atoi(verticesEdges[0])

	//edges := verticesEdges[1]
	//_ = edges

	input := lines[1:]
	arrInput := make([][]string, vertices)

	for _, v := range input {
		if v == "" {
			continue
		}

		tmp := strings.Split(strings.Replace(v, "\r", "", -1), " ")
		edgeOut, _ := strconv.Atoi(tmp[0])
		edgeIn, _ := strconv.Atoi(tmp[1])

		if arrInput[edgeOut-1] == nil {
			arrInput[edgeOut-1] = makeArrString(vertices)
		}

		arrInput[edgeOut-1][edgeIn-1] = "1"
	}

	result := make([]string, vertices)
	for k, v := range arrInput {
		if v == nil {
			v = makeArrString(vertices)
		}
		result[k] = strings.Join(v, " ") + " "
	}

	return strings.Join(result, "\n") + "\n"
}
