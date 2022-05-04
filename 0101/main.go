package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

func helper(t1 *TreeNode, t2 *TreeNode) bool {

	if t1 == nil && t2 == nil {
		return true
	}
	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
		return false
	}
	return t1.Val == t2.Val && helper(t1.Left, t2.Right) && helper(t1.Right, t2.Left)
}
