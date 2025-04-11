package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * Solution: Recursion with DFS
 */
func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	var fn func(root *TreeNode) int
	fn = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := fn(root.Left)
		right := fn(root.Right)
		ret = max(ret, left+right)
		return 1 + max(left, right)
	}
	fn(root)
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
