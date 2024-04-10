package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * Using two poniters to solve this problem
 * 1. Create a dummy node to handle the edge case
 * 2. Move right pointer to n steps
 * 3. Move left and right pointers to end
 * 4. Remove the node
 * 5. Return dummy.Next
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	l, r := dummy, dummy
	// move right
	for i := 0; i < n; i++ {
		r = r.Next
	}
	// move right to end
	for r.Next != nil {
		r = r.Next
		l = l.Next
	}
	// remove node
	l.Next = l.Next.Next

	return dummy.Next
}
