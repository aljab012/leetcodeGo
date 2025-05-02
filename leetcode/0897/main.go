package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	prev := dummy

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		node.Left = nil
		prev.Right = node
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	return dummy.Right
}
