package main

import (
	"strconv"
	"strings"
)

type Node struct {
	value contestant
	left  *Node
	right *Node
}

type contestant struct {
	name    string
	solved  int
	penalty int
}

func heapify(arr []contestant, n, i int) {
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && !ASmallerB(arr[i], arr[left]) {
			largest = left
		}

		if right < n && !ASmallerB(arr[largest], arr[right]) {
			largest = right
		}

		if largest == i {
			break
		}

		arr[i], arr[largest] = arr[largest], arr[i]
		i = largest
	}
}

func buildHeap(arr []contestant) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func sortHeap(arr []contestant) {
	n := len(arr)

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func ASmallerB(a, b contestant) bool {
	if a.solved != b.solved {
		return a.solved < b.solved
	}
	if a.penalty != b.penalty {
		return a.penalty > b.penalty
	}
	return a.name > b.name
}

func finalA(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")

	commands := lines[1:]
	heap := make([]contestant, len(commands))

	for k, v := range commands {
		tmp := strings.Split(strings.Replace(v, "\r", "", -1), " ")

		solved, _ := strconv.Atoi(tmp[1])
		penalty, _ := strconv.Atoi(tmp[2])

		heap[k] = contestant{name: tmp[0], solved: solved, penalty: penalty}
		_ = k
	}

	buildHeap(heap)
	sortHeap(heap)

	result := make([]string, len(heap))
	for k, v := range heap {
		result[k] = v.name
	}

	return strings.Join(result, "\n") + "\n"
}
