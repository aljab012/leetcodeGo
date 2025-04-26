package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	ret := 0

	var visit func(node *TreeNode) (sum int, count int)
	visit = func(node *TreeNode) (sum int, count int) {
		if node == nil {
			return 0, 0
		}

		leftSum, leftCount := visit(node.Left)
		rightSum, rightCount := visit(node.Right)

		totalSum := node.Val + leftSum + rightSum
		totalCount := 1 + leftCount + rightCount

		if (totalSum / totalCount) == node.Val {
			ret++
		}

		return totalSum, totalCount
	}

	visit(root)

	return ret
}
