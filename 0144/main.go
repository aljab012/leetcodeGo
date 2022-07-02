package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return helper(root, []int{})
}
func helper(root *TreeNode, ret []int) []int {
	if root == nil {
		return ret
	}
	ret = append(ret, root.Val)
	ret = helper(root.Left, ret)
	return helper(root.Right, ret)
}
