package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var fn func(root *TreeNode) int
	fn = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + max(fn(root.Left), fn(root.Right))
	}
	return fn(root)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
