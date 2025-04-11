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
	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := []int{}
		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]

			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
			if len(ret)%2 == 0 {
				levelValues = append(levelValues, pop.Val)
			} else {
				levelValues = append([]int{pop.Val}, levelValues...)
			}
		}
		ret = append(ret, levelValues)
	}

	return ret
}
