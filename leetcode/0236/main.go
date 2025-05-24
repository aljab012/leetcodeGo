package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	leftSubtree := lowestCommonAncestor(root.Left, p, q)
	rightSubtree := lowestCommonAncestor(root.Right, p, q)
	if leftSubtree != nil && rightSubtree != nil {
		return root
	}

	if leftSubtree != nil {
		return leftSubtree
	}
	return rightSubtree
}
