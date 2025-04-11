package main

func isMonotonic(nums []int) bool {
	inc, dec := true, true

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dec = false
		}
		if nums[i] < nums[i-1] {
			inc = false
		}
	}
	return inc || dec
}
