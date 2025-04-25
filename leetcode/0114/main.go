package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var last *TreeNode
	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		visit(node.Right)
		visit(node.Left)
		node.Right = last
		node.Left = nil
		last = node
	}
	visit(root)
}
