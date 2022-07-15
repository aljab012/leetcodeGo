package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func invertTree2(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pop.Left, pop.Right = pop.Right, pop.Left

		if pop.Left != nil {
			stack = append(stack, pop.Left)
		}

		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}
	}
	return root
}
