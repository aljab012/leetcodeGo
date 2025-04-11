package main

// 1 2 3
func nextPermutation(nums []int) {

	// find pivot
	pivot := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			pivot = i
			break
		}
	}
	// swap pivot
	if pivot >= 0 {
		for i := len(nums) - 1; i > pivot; i-- {
			if nums[i] > nums[pivot] {
				swap(nums, i, pivot)
				break
			}
		}
	}
	// reverse the right side
	reverse(nums, pivot+1, len(nums)-1)
}
func reverse(nums []int, i int, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, x int, y int) {
	temp := nums[x]
	nums[x] = nums[y]
	nums[y] = temp
}

func main() {
	nextPermutation([]int{1, 2, 3})
}
