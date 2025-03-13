package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ret := []int{}

	var fn func(root *TreeNode)
	fn = func(root *TreeNode) {
		if root == nil {
			return
		}
		ret = append(ret, root.Val)
		fn(root.Left)
		fn(root.Right)
	}

	fn(root)
	return ret
}
