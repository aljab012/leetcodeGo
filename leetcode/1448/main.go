package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	count := 0

	var visit func(node *TreeNode, maxSoFar int)
	visit = func(node *TreeNode, maxSoFar int) {
		if node == nil {
			return
		}
		if node.Val >= maxSoFar {
			count++
		}
		visit(node.Left, max(node.Val, maxSoFar))
		visit(node.Right, max(node.Val, maxSoFar))
	}

	visit(root, math.MinInt)
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
