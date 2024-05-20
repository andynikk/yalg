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

func dfs(arr [][]int, start int) string {

	color := make([]int, len(arr))
	times := make([]string, len(arr))
	time := 0
	stack := []int{start}

	for len(stack) > 0 {
		apex := stack[len(stack)-1]
		stack = (stack)[:len(stack)-1]

		if color[apex] == white {
			times[apex] = strconv.Itoa(time)
			color[apex] = grey
			time++
			stack = append(stack, apex)

			for _, adjacent := range arr[apex] {
				if color[adjacent] == white {
					stack = append(stack, adjacent)
				}
			}
		} else if color[apex] == grey {
			times[apex] = times[apex] + " " + strconv.Itoa(time)
			color[apex] = black
			time++
		}
	}

	if times[0] == "" {
		times = times[1:]
	}

	return strings.Join(times, "\n") + "\n"
}

func task08(res *[]byte) string {
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
	}
	for i := 0; i < len(arrInput); i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(arrInput[i])))
	}

	//times := dfs(arrInput, 1)
	//
	//result := []string{}
	//for i := 0; i < len(times); i++ {
	//	if times[i][0] == "" && times[i][1] == "" {
	//		continue
	//	}
	//
	//	result = append(result, times[i][0]+" "+times[i][1])
	//}

	return dfs(arrInput, 1)
}
