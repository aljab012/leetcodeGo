package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	startK := len(nums) - k
	temp := append(append([]int{}, nums[startK:]...), nums[:startK]...)
	copy(nums, temp)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
