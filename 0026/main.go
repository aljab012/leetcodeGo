package main

import "fmt"

// idea: Have pivot index that only moves with unique values from the begining. Iterate through the array,
// if same as pivot index, skip. Otherwise, insert into a pivot and move the pivot index.
func removeDuplicates(nums []int) int {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[index] {
			continue
		}
		index += 1
		nums[index] = nums[i]
	}
	return index + 1

}

func main() {
	nums := []int{1, 1, 2}
	ret := removeDuplicates(nums)
	fmt.Println(ret)
}
