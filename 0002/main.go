package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := new(ListNode)
	carry := 0

	for cur := dummy; l1 != nil || l2 != nil || carry > 0; cur = cur.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{carry % 10, nil}
		carry /= 10

	}
	return dummy.Next

}
