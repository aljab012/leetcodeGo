package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// naive solution using a set
func hasCycle1(head *ListNode) bool {
	set := map[*ListNode]bool{}
	cur := head

	for cur != nil {
		if _, ok := set[cur]; ok {
			return true
		}
		set[cur] = true
		cur = cur.Next
	}
	return false
}

// Floyd's cycle detection algorithm
func hasCycle2(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
