package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return helper(0, nums, []int{}, [][]int{})
}
func helper(start int, nums []int, cur []int, res [][]int) [][]int {
	res = append(res, append([]int{}, cur...))
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		res = helper(i+1, nums, append(cur, nums[i]), res)
	}
	return res
}
