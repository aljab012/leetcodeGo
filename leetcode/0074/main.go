package main

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	for l <= r {
		mid := l + (r-l)/2
		if target < matrix[mid][0] {
			r = mid - 1
		} else if target > matrix[mid][len(matrix[mid])-1] {
			l = mid + 1
		} else {
			return binarySearch(matrix[mid], target)
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}
