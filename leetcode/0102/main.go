package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		qLen := len(queue)
		level := []int{}
		for i := 0; i < qLen; i++ {
			element := queue[0]
			queue = queue[1:]

			level = append(level, element.Val)
			if element.Left != nil {
				queue = append(queue, element.Left)
			}
			if element.Right != nil {
				queue = append(queue, element.Right)
			}
		}
		ret = append(ret, level)
	}
	return ret
}
