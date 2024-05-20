package sprint2

// <template>
type ListNode3 struct {
	data string
	next *ListNode3
}

// <template>

func Solution3(head *ListNode3, idx int) *ListNode3 {
	// Your code
	// ヽ(´▽`)/

	predHead := &ListNode3{}
	currHead := head

	if idx == 0 {
		*head = *head.next
		return head
	}

	for i := 0; i <= idx; i++ {

		if i != idx {
			predHead = currHead
			currHead = currHead.next

			continue
		}

		predHead.next = currHead.next
	}

	return head
}

func test3() {
	node3 := ListNode3{"node3", nil}
	node2 := ListNode3{"node2", &node3}
	node1 := ListNode3{"node1", &node2}
	node0 := ListNode3{"node0", &node1}
	/*newHead :=*/ Solution3(&node0, 0)
	// result is : node0 -> node2 -> node3
}

func Task3() {
	test3()
}
