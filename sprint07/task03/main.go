package main

import (
	"sort"
	"strconv"
	"strings"
)

func fifo(descriptionHeaps [][2]int, maximum int) string {

	toBeDebited := 0
	remains := maximum
	sum := 0

	for _, heap := range descriptionHeaps {
		if remains <= 0 {
			break
		}

		toBeDebited = min(remains, heap[1])
		sum += toBeDebited * heap[0]

		remains -= toBeDebited
	}

	return strconv.Itoa(sum) + "\n"
}

// <template>
func goldRush(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	tonnage, _ := strconv.Atoi(lines[0])
	heaps, _ := strconv.Atoi(lines[1])

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	lines = lines[2:]

	descriptionHeaps := make([][2]int, heaps)
	for i := 0; i < len(lines); i++ {
		line := strings.Split(strings.Replace(lines[i], "\r", "", -1), " ")

		kg, _ := strconv.Atoi(line[0])
		price, _ := strconv.Atoi(line[1])

		descriptionHeaps[i] = [2]int{kg, price}
	}
	sort.Slice(descriptionHeaps, func(i, j int) bool {
		return descriptionHeaps[i][0] > descriptionHeaps[j][0]
	})

	result := fifo(descriptionHeaps, tonnage)

	return result
}
