package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}

	var dfs func(node *TreeNode, pathSum int, path []int)
	dfs = func(node *TreeNode, pathSum int, path []int) {
		if node == nil {
			return
		}
		pathSum += node.Val
		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil && pathSum == targetSum {
			result = append(result, append([]int{}, path...))
		}

		dfs(node.Left, pathSum, path)
		dfs(node.Right, pathSum, path)
	}

	dfs(root, 0, []int{})
	return result
}
