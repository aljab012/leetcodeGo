package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 1
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]

			if pop.Left == nil && pop.Right == nil {
				return depth
			}

			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		depth++
	}
	return depth
}
