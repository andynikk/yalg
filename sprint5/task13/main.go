package main

func siftUp(heap []int, idx int) int {
	// Your code
	// “ヽ(´▽｀)ノ”

	for idx > 1 {
		parentIdx := idx / 2

		if heap[idx] > heap[parentIdx] {
			heap[idx], heap[parentIdx] = heap[parentIdx], heap[idx]
			idx = parentIdx
		} else {
			break
		}
	}

	return idx

}

func test() {
	sample := []int{-1, 12, 6, 8, 3, 15, 7}
	if siftUp(sample, 5) != 1 {
		panic("WA")
	}
}
