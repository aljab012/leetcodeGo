package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * 1. If both trees are nil, they are the same.
 * 2. If one tree is nil and the other is not, they are not the same.
 * 3. If both trees are not nil, check if the values of the current nodes are
 *    equal and check the left and right subtrees recursively.
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
