package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var fn func(root *TreeNode)
	fn = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val >= low && root.Val <= high {
			sum += root.Val
		}
		fn(root.Left)
		fn(root.Right)
	}
	fn(root)
	return sum
}
