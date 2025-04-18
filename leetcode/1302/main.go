package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return -1
	}

	queue := []*TreeNode{root}
	lastLevelSum := 0
	for len(queue) > 0 {
		lastLevelSum = 0
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]
			lastLevelSum += pop.Val
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
	}

	return lastLevelSum
}
