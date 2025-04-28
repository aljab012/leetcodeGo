package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSameTree(root, subRoot) ||
		isSubtree(root.Left, subRoot) ||
		isSubtree(root.Right, subRoot)
}
func isSameTree(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val &&
		isSameTree(node1.Left, node2.Left) &&
		isSameTree(node1.Right, node2.Right)
}
