package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	queue := []*Node{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]

			for _, child := range pop.Children {
				queue = append(queue, child)
			}
		}
		depth++
	}
	return depth
}
