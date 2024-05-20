package sprint2

// <template>
type ListNode2 struct {
	data string
	next *ListNode2
}

// <template>

func Solution2(head *ListNode2) {
	println(head.data)

	if head.next == nil {
		return
	}
	Solution2(head.next)

}

func test2() {
	node3 := ListNode2{"node3", nil}
	node2 := ListNode2{"node2", &node3}
	node1 := ListNode2{"node1", &node2}
	node0 := ListNode2{"node0", &node1}
	Solution2(&node0)
	/*
	   Output is:
	   node0
	   node1
	   node2
	   node3
	*/
}

func Task2() {
	test2()
}
