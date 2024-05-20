package main

import (
	"strconv"
	"strings"
)

func task01(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")

	verticesEdges := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")
	vertices, _ := strconv.Atoi(verticesEdges[0])

	input := lines[1:]
	arrInput := make([][]string, vertices)

	for _, v := range input {
		if v == "" {
			continue
		}

		tmp := strings.Split(strings.Replace(v, "\r", "", -1), " ")
		edgeOut, _ := strconv.Atoi(tmp[0])
		//edgeIn, _ := strconv.Atoi(tmp[1])

		if arrInput[edgeOut-1] == nil {
			arrInput[edgeOut-1] = []string{"1", tmp[1]}
		} else {
			r, _ := strconv.Atoi(arrInput[edgeOut-1][0])
			arrInput[edgeOut-1][0] = strconv.Itoa(r + 1)
			arrInput[edgeOut-1] = append(arrInput[edgeOut-1], tmp[1])
		}
	}

	result := make([]string, vertices)
	for k, v := range arrInput {
		if v == nil {
			result[k] = "0 "
			continue
		}

		result[k] = strings.Join(v, " ") + " "
	}

	return strings.Join(result, "\n") + "\n"
}
