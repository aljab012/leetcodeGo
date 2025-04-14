package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* Recursive
 * 1. If both trees are nil, they are the same.
 * 2. If one tree is nil and the other is not, they are not the same.
 * 3. If both trees are not nil, check if the values of the current nodes are
 *    equal and check the left and right subtrees recursively.
 */
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree1(p.Left, q.Left) && isSameTree1(p.Right, q.Right)
}

/*
 * Iterative
 * 1. Use a queue to store the nodes of both trees.
 * 2. While the queue is not empty, pop two nodes from the queue.
 * 3. If both nodes are nil, continue to the next iteration.
 * 4. If one node is nil and the other is not, return false.
 * 5. If both nodes are not nil, check if their values are equal.
 * 6. Add the left and right children of both nodes to the queue.
 * 7. If all nodes are processed and no differences are found, return true.
 */
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}

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
		queue = append(queue, node2.Left)
		queue = append(queue, node1.Right)
		queue = append(queue, node2.Right)
	}
	return true
}
