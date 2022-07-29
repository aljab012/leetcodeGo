package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	return helper(root, []int{})
}

func helper(root *Node, ret []int) []int {
	if root == nil {
		return ret
	}
	ret = append(ret, root.Val)
	for i := range root.Children {
		ret = helper(root.Children[i], ret)
	}
	return ret
}
