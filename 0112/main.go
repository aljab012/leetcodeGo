package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	return helper(root, targetSum)
}

func helper(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}
	sum := targetSum - root.Val
	if sum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return helper(root.Left, sum) || helper(root.Right, sum)
}
