package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		// set next pointer
		for i := 0; i < levelSize-1; i++ {
			queue[i].Next = queue[i+1]
		}

		// pop node and add children
		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
	}
	return root
}
