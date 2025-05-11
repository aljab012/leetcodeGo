package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	result := 0
	index := 1

	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}

		visit(node.Left)
		if index == k {
			result = node.Val
		}
		index++
		visit(node.Right)
	}

	visit(root)
	return result
}
