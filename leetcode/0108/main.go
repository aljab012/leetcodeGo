package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIndex := len(nums) / 2
	root := &TreeNode{Val: nums[midIndex]}
	root.Left = sortedArrayToBST(nums[:midIndex])
	root.Right = sortedArrayToBST(nums[midIndex+1:])
	return root
}
