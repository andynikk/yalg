package sprint2

// <template>
type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

// <template>

func (l *ListNode) Revers() {

	tmpPrev := l.prev

	l.prev = l.next
	l.next = tmpPrev

	if l.next == nil {
		return
	}

	l.next.Revers()
}

func Solution(head *ListNode) *ListNode {
	// Your code
	// ヽ(´▽`)/

	if head.next == nil {
		return head
	}

	newHead := head.next
	for {
		if newHead.next == nil {
			break
		}

		newHead = newHead.next
	}

	newHead.Revers()

	return newHead

}

func test5() {
	node3 := ListNode{"node3", nil, nil}
	node2 := ListNode{"node2", &node3, nil}
	node1 := ListNode{"node1", &node2, nil}
	node0 := ListNode{"node0", &node1, nil}
	node3.prev = &node2
	node2.prev = &node1
	node1.prev = &node0
	/*newHead :=*/ Solution(&node0)
	// result is : newHead == node3
	// node3.next == node2
	// node2.next == node1
	// node2.prev = node3
	// node1.next == node0
	// node1.prev == node2
	// node0.prev == node1
}

func Task5() {
	test5()
}
