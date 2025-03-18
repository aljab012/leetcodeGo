package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	minVale := math.MinInt64
	isValid := true

	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		visit(node.Left)
		if node.Val <= minVale {
			isValid = false
		}
		minVale = node.Val
		visit(node.Right)
	}
	visit(root)
	return isValid
}
