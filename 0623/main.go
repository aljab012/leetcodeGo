package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &TreeNode{Left: root, Val: val}
	}
	type NodeWithDepth struct {
		node  *TreeNode
		depth int
	}
	queue := []NodeWithDepth{{node: root, depth: 1}}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]

			if pop.depth == depth-1 {
				pop.node.Left = &TreeNode{Val: val, Left: pop.node.Left}
				pop.node.Right = &TreeNode{Val: val, Right: pop.node.Right}
			}

			if pop.node.Left != nil {
				queue = append(queue, NodeWithDepth{node: pop.node.Left, depth: pop.depth + 1})
			}
			if pop.node.Right != nil {
				queue = append(queue, NodeWithDepth{node: pop.node.Right, depth: pop.depth + 1})
			}
		}
	}
	return root
}
