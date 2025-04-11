package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 1. Find the middle of the list
 * 2. Reverse the second half
 * 3. Merge the two halves
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */
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
