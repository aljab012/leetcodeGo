package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	result := []int{}

	nodeCount := 0
	frequency := map[int]int{}

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		nodeCount++
		frequency[sum]++

		return sum
	}
	dfs(root)

	for len(result) == 0 {
		for key, count := range frequency {
			if count == nodeCount {
				result = append(result, key)
			}
		}
		nodeCount--
	}
	return result
}
