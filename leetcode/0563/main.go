package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	result := 0

	var treeSum func(node *TreeNode) int
	treeSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := treeSum(node.Left)
		rightSum := treeSum(node.Right)

		result += abs(leftSum - rightSum)

		return node.Val + leftSum + rightSum
	}

	treeSum(root)
	return result
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
