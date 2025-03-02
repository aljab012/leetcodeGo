package main

func searchInsert(nums []int, target int) int {
	isBefore := func(i int) bool {
		return nums[i] <= target
	}

	left, right := 0, len(nums)-1
	// handle edge cases
	if target < nums[left] {
		return 0
	}
	if target > nums[right] {
		return right + 1
	}

	for (right - left) > 1 {
		mid := (left + right) / 2
		if isBefore(mid) {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[left] < target {
		return left + 1
	}
	return left
}
