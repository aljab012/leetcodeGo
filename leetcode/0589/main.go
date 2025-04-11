package main

type Node struct {
	Val      int
	Children []*Node
}

/*
 * Recursive
 */
func preorder1(root *Node) []int {
	ret := []int{}

	var visit func(node *Node)
	visit = func(node *Node) {
		if node == nil {
			return
		}
		ret = append(ret, node.Val)
		for _, child := range node.Children {
			visit(child)
		}
	}

	visit(root)

	return ret
}

/*
 * Iterative
 */

func preorder2(root *Node) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	stack := []*Node{root}
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ret = append(ret, pop.Val)
		for i := len(pop.Children) - 1; i >= 0; i-- {
			stack = append(stack, pop.Children[i])
		}
	}
	return ret
}
