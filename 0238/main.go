package main

/*
 * solution with calculating the prefix and postfix
 * Space complexity: O(1)
 * Time complexity: O(n)
 */
func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	prefix := 1
	postfix := 1
	for i := range nums {
		ret[i] = prefix
		prefix *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= postfix
		postfix *= nums[i]

	}
	return ret
}
