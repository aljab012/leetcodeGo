package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	isOddLevel := true
	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]

			level = append(level, pop.Val)

			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		// validate numbers
		for i, n := range level {
			if isOddLevel {
				if n%2 == 0 {
					return false
				}
				if i != 0 && level[i] <= level[i-1] {
					return false
				}
			} else {
				if n%2 == 1 {
					return false
				}
				if i != 0 && level[i] >= level[i-1] {
					return false
				}
			}
		}
		isOddLevel = !isOddLevel
	}
	return true
}
