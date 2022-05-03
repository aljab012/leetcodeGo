package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// idea: base case: if one of them if null, return the other. otherwise, compare the values. merge the later and next in a recursive call.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2

}

func merge()
