package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	ret := []int{}

	if root == nil {
		return ret
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		levelLen := len(stack)
		max := math.MinInt
		for i := 0; i < levelLen; i++ {
			pop := stack[0]
			stack = stack[1:]
			if pop.Val > max {
				max = pop.Val
			}
			if pop.Left != nil {
				stack = append(stack, pop.Left)
			}
			if pop.Right != nil {
				stack = append(stack, pop.Right)
			}
		}
		ret = append(ret, max)
	}

	return ret
}
