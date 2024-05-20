package main

func siftDown(heap []int, idx int) int {
	// Your code
	// “ヽ(´▽｀)ノ”

	n := len(heap) - 1
	for {
		leftChildIdx := 2 * idx
		rightChildIdx := 2*idx + 1
		largest := idx

		if leftChildIdx <= n && heap[leftChildIdx] > heap[idx] {
			largest = leftChildIdx
		}

		if rightChildIdx <= n && heap[rightChildIdx] > heap[largest] {
			largest = rightChildIdx
		}

		if largest == idx {
			break
		}

		heap[idx], heap[largest] = heap[largest], heap[idx]
		idx = largest
	}

	return idx

}

func test() {
	sample := []int{-1, 12, 1, 8, 3, 4, 7}
	if siftDown(sample, 2) != 5 {
		panic("WA")
	}
}
