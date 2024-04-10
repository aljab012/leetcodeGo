package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * Using two poniters to solve this problem interatively
 * Save the next node in a variable
 * Reverse the current node
 * Move to the next node
 * Time complexity: O(n)
 * Space complexity: O(1)
 */

func reverseList(head *ListNode) *ListNode {

	var prev, cur *ListNode = nil, head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
