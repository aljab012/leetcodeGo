package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	seenNil := false
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]

			if pop == nil {
				seenNil = true
				continue
			}
			if seenNil {
				return false
			}
			queue = append(queue, pop.Left)
			queue = append(queue, pop.Right)
		}
	}
	return true
}
