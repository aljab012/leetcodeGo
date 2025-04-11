package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		levelMax := queue[0].Val

		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]

			if pop.Val > levelMax {
				levelMax = pop.Val
			}

			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		ret = append(ret, levelMax)
	}
	return ret
}
