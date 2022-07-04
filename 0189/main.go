package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	nums = reverse(nums, 0, len(nums)-1)
	nums = reverse(nums, 0, k-1)
	nums = reverse(nums, k, len(nums)-1)
}

func reverse(arr []int, l, r int) []int {
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return arr
}

func main() {
	fmt.Print()
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
