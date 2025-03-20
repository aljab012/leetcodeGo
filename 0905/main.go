package main

func sortArrayByParity(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]%2 == 0 { // left is even, skip
			l++
		} else if nums[r]%2 == 1 { // right is odd, skip
			r--
		} else { // swap left and right
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return nums
}
