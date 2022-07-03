package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	temp := nums
	for i := 0; i < k; i++ {
		temp = rotateOne(temp)
	}
	copy(nums, temp)
}
func rotateOne(nums []int) []int {
	last := nums[len(nums)-1]
	nums = append([]int{last}, nums[:len(nums)-1]...)
	return nums
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
