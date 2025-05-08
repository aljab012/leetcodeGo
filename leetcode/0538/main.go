package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		visit(node.Right)
		sum += node.Val
		node.Val = sum
		visit(node.Left)
	}

	visit(root)
	return root
}
