package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := inOrderTraversal(root1)
	arr2 := inOrderTraversal(root2)

	return mergeSortedArray(arr1, arr2)
}

func mergeSortedArray(arr1, arr2 []int) []int {
	ret := []int{}
	p1, p2 := 0, 0

	for p1 < len(arr1) && p2 < len(arr2) {
		if arr1[p1] < arr2[p2] {
			ret = append(ret, arr1[p1])
			p1++
		} else {
			ret = append(ret, arr2[p2])
			p2++
		}
	}
	ret = append(ret, arr1[p1:]...)
	ret = append(ret, arr2[p2:]...)

	return ret
}

func inOrderTraversal(root *TreeNode) []int {
	result := []int{}

	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		visit(node.Left)
		result = append(result, node.Val)
		visit(node.Right)
	}

	visit(root)
	return result
}
