package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	isBalanced := true
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)
		if abs(leftDepth-rightDepth) > 1 {
			isBalanced = false
		}
		return 1 + max(leftDepth, rightDepth)
	}

	depth(root)
	return isBalanced
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
