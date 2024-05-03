package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * Recursive approach
 * Time complexity: O(n)
 * Space complexity: O(n)
 * Intuition: We can invert a binary tree by recursively inverting its left and right subtrees.
 */
func invertTree1(root *TreeNode) *TreeNode {
	var fn func(r *TreeNode) *TreeNode
	fn = func(r *TreeNode) *TreeNode {
		if r == nil {
			return nil
		}
		r.Left, r.Right = fn(r.Right), fn(r.Left)
		return r
	}
	return fn(root)
}

/*
 * Iterative approach
 * Time complexity: O(n)
 * Space complexity: O(n)
 * Intuition: We can also solve this problem iteratively. We start by creating a queue of nodes to visit.
 * We then visit each node in the queue, swapping its left and right children.
 * If the left child is not nil, we add it to the queue.
 * If the right child is not nil, we add it to the queue.
 * We continue this process until the queue is empty.
 */
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	arr := []*TreeNode{root}

	for len(arr) != 0 {
		top := arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		top.Left, top.Right = top.Right, top.Left
		if top.Left != nil {
			arr = append(arr, top.Left)
		}
		if top.Right != nil {
			arr = append(arr, top.Right)
		}
	}
	return root

}
