package main

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}
	return result
}
