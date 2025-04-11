package main

// idea: Have pivot index that only moves with unique values from the begining. Iterate through the array,
// if same as pivot index, skip. Otherwise, insert into a pivot and move the pivot index.
func removeDuplicates(nums []int) int {
	pivot := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[pivot] {
			continue
		}
		pivot++
		nums[pivot] = nums[i]
	}
	return pivot + 1
}

// shorter
func removeDuplicates2(nums []int) int {
	pivot := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[pivot] {
			pivot++
			nums[pivot] = nums[i]
		}
	}
	return pivot + 1
}
