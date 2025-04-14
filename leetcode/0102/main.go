package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * BFS
 * 1. If the root is nil, return an empty slice.
 * 2. Initialize a queue with the root node.
 * 3. While the queue is not empty, do the following:
 *  a. Get the length of the queue.
 *  b. Initialize an empty slice for the current level.
 *  c. For each node in the queue, do the following:
 *  a. Pop the first node from the queue.
 *  b. Append its value to the current level slice.
 *  c. If the left child is not nil, append it to the queue.
 *  d. If the right child is not nil, append it to the queue.
 *  4. Append the current level slice to the result slice.
 *  5. Return the result slice.
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		qLen := len(queue)
		level := []int{}
		for i := 0; i < qLen; i++ {
			element := queue[0]
			queue = queue[1:]

			level = append(level, element.Val)
			if element.Left != nil {
				queue = append(queue, element.Left)
			}
			if element.Right != nil {
				queue = append(queue, element.Right)
			}
		}
		ret = append(ret, level)
	}
	return ret
}
