package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	isOdd := false

	for len(queue) > 0 {
		levelSize := len(queue)
		level := queue[:levelSize]
		queue = queue[levelSize:]

		// Add children to queue for next level
		for _, node := range level {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Reverse values for odd levels
		if isOdd {
			l, r := 0, len(level)-1
			for l < r {
				level[l].Val, level[r].Val = level[r].Val, level[l].Val
				l++
				r--
			}
		}
		isOdd = !isOdd
	}
	return root
}
