package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, index := m-1, n-1, m+n-1
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[index] = nums1[p1]
			p1--
		} else {
			nums1[index] = nums2[p2]
			p2--
		}
		index--
	}
}
