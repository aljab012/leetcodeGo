package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	lastSum := 0
	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		visit(node.Right)
		lastSum += node.Val
		node.Val = lastSum
		visit(node.Left)
	}
	visit(root)
	return root
}
