package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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

			levelValues = append(levelValues, pop.Val)
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		ret = append([][]int{levelValues}, ret...)
	}
	return ret
}
