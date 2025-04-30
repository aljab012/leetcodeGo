package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	nMap := map[int]bool{}

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		complement := k - node.Val
		if nMap[complement] {
			return true
		}
		nMap[node.Val] = true

		return dfs(node.Left) || dfs(node.Right)

	}

	return dfs(root)
}
