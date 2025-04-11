package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ret := [][]int{}

	if root == nil {
		return ret
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := []int{}

		for i := 0; i < levelSize; i++ {
			pop := queue[0]
			queue = queue[1:]

			levelValues = append(levelValues, pop.Val)
			for _, node := range pop.Children {
				queue = append(queue, node)
			}
		}
		ret = append(ret, levelValues)
	}
	return ret
}
