package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	// get length and tail
	length := 1
	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	k = k % length
	if k == 0 {
		return head
	}

	// iterate and split
	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	tail.Next = head

	return newHead
}
