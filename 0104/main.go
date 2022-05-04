package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return helper(root)

}
func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(helper(root.Left), helper(root.Right))
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
