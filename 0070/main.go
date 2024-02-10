package main

// space complexity: O(1)
// time complexity: O(n)
func climbStairs(n int) int {
	prevStep, curStep := 1, 1
	for i := 2; i <= n; i++ {
		curStep, prevStep = prevStep+curStep, curStep
	}
	return curStep
}
