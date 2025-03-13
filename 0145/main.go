package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ret := []int{}

	var fn func(root *TreeNode)
	fn = func(root *TreeNode) {
		if root == nil {
			return
		}
		fn(root.Left)
		fn(root.Right)
		ret = append(ret, root.Val)
	}

	fn(root)
	return ret
}
