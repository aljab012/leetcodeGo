package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}
	isReversed := false
	for len(queue) != 0 {
		qLen := len(queue)
		level := []int{}
		for i := 0; i < qLen; i++ {
			element := queue[0]
			queue = queue[1:]
			if element.Left != nil {
				queue = append(queue, element.Left)
			}
			if element.Right != nil {
				queue = append(queue, element.Right)
			}
			if isReversed {
				level = append([]int{element.Val}, level...)
			} else {
				level = append(level, element.Val)
			}
		}
		isReversed = !isReversed
		ret = append(ret, level)
	}
	return ret
}
