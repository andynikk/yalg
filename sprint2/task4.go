package sprint2

// <template>
type ListNode4 struct {
	data string
	next *ListNode4
}

// <template>

func Solution4(head *ListNode4, elem string) int {
	// Your code
	// ヽ(´▽`)/

	curNode := head
	i := -1
	for {
		i++

		if curNode.next == nil {
			break
		}

		if curNode.data != elem {
			curNode = curNode.next
			continue
		}

		return i
	}

	return -1
}

func test4() {
	node3 := ListNode4{"node3", nil}
	node2 := ListNode4{"node2", &node3}
	node1 := ListNode4{"node1", &node2}
	node0 := ListNode4{"node0", &node1}
	/*idx :=*/ Solution4(&node0, "node2")
	// result is : idx == 2
}

func Task4() {
	test4()
}
