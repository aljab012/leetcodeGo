package main

/*
 * solution with calculating the prefix and postfix
 * Space complexity: O(1)
 * Time complexity: O(n)
 */
func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	prefix := 1
	for i, n := range nums {
		ret[i] = prefix
		prefix *= n
	}
	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= suffix
		suffix *= nums[i]
	}
	return ret
}
