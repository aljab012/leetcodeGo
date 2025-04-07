package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return -1
	}

	queue := []*TreeNode{root}
	curLevel := 1

	maxSum := root.Val
	maxLevel := 1

	for len(queue) > 0 {
		levelSize := len(queue)
		curLevelSum := 0

		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]
			curLevelSum += pop.Val

			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}

		if curLevelSum > maxSum {
			maxSum = curLevelSum
			maxLevel = curLevel
		}
		curLevel++
	}

	return maxLevel
}
