package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
 * Idea: create a map to store the old node to new node mapping
 * iterate through the list to create the new list
 * then iterate through the list to set the random pointer and next pointer
 * Space complexity: O(n)
 * Time complexity: O(n)
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	oldToNewMap := map[*Node]*Node{}

	// iterate through the list to create the new list
	cur := head
	for cur != nil {
		oldToNewMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	// iterate through the list to set the random pointer and next pointer
	cur = head
	for cur != nil {
		newNode := oldToNewMap[cur]
		newNode.Next = oldToNewMap[cur.Next]
		newNode.Random = oldToNewMap[cur.Random]
		cur = cur.Next
	}
	return oldToNewMap[head]
}
