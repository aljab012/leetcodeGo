package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// reverse second half
	var prev *ListNode
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}

	// merge
	first := head
	for prev.Next != nil {
		first.Next, first = prev, first.Next
		prev.Next, prev = first, prev.Next
	}
}
