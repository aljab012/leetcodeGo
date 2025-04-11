package main

/*
 * Approach:
 * ToolBox: Two Pointers
 * 1. Use two pointers to traverse the two arrays from the end.
 * 2. Compare the elements and place the larger one at the end of nums1.
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p2 >= 0 && p1 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p--
			p1--
		} else {
			nums1[p] = nums2[p2]
			p--
			p2--
		}
	}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p--
		p2--
	}
}
