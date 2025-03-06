package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 2. Add Two Numbers
 * Intuition: We can add two numbers by iterating through both lists and adding the values of the nodes along with the carry.
 * Time complexity: O(max(m, n)) where m and n are the lengths of the two input lists.
 * Space complexity: O(max(m, n)) where m and n are the lengths of the two input lists.
 */
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 += l2.Val
			l2 = l2.Next
		}
		sum := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}

// simpler version
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

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
