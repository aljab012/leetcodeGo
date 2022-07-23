package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	l1 = reverse(l1)
	l2 = reverse(l2)
	carry := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		newNode := &ListNode{Val: carry % 10}
		cur.Next = newNode
		cur = cur.Next
		carry = carry / 10
	}
	if carry != 0 {
		newNode := &ListNode{Val: carry}
		cur.Next = newNode
		cur = cur.Next
	}
	return reverse(dummyHead.Next)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
