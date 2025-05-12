package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	minDiff := math.MaxInt

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil {
			minDiff = min(minDiff, node.Val-prev.Val)
		}
		prev = node
		dfs(node.Right)
	}

	dfs(root)
	return minDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
