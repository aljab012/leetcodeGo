package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	var fn func(node *TreeNode, sum int) int
	fn = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}
		sum = sum<<1 | node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return fn(node.Left, sum) + fn(node.Right, sum)
	}
	return fn(root, 0)
}
