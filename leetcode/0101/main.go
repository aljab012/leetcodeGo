package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * Recursive solution
 * 1. If both trees are nil, they are the same.
 * 2. If one tree is nil and the other is not, they are not the same.
 * 3. If both trees are not nil, check if the values of the current nodes are
 *   equal and check the left and right subtrees recursively.
 */
func isSymmetric1(root *TreeNode) bool {
	var fn func(node1, node2 *TreeNode) bool
	fn = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}
		return node1.Val == node2.Val && fn(node1.Left, node2.Right) && fn(node1.Right, node2.Left)
	}
	return fn(root, root)
}

/**
 * Iterative solution
 * 1. Use a queue to store the nodes of both trees.
 * 2. While the queue is not empty, pop two nodes from the queue.
 * 3. If both nodes are nil, continue to the next iteration.
 * 4. If one node is nil and the other is not, return false.
 * 5. If both nodes are not nil, check if their values are equal.
 * 6. Add the left and right children of both nodes to the queue.
 * 7. If all nodes are processed and no differences are found, return true.
 * 8. The tree is symmetric if the left and right subtrees are mirror images of each other.
 */
func isSymmetric2(root *TreeNode) bool {
	queue := []*TreeNode{root, root}

	for len(queue) > 0 {
		node1, node2 := queue[0], queue[1]
		queue = queue[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		queue = append(queue, node1.Left)
		queue = append(queue, node2.Right)
		queue = append(queue, node1.Right)
		queue = append(queue, node2.Left)
	}
	return true
}
