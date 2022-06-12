package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// idea: base case: if one of them if null, return the other. otherwise, compare the values. merge the later and next in a recursive call.
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)

	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{l1.Val, nil}
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{l2.Val, nil}
			cur = cur.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		cur.Next = &ListNode{l1.Val, nil}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		cur.Next = &ListNode{l2.Val, nil}
		cur = cur.Next
		l2 = l2.Next
	}
	return dummy.Next
}
