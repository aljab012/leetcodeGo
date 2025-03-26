package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ret := []int{}

	var visit func(node *Node)
	visit = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			visit(child)
		}
		ret = append(ret, node.Val)
	}

	visit(root)
	return ret
}
