package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	p1 := m - 1
	p2 := n - 1
	index := m + n - 1

	for p1 >= 0 && p2 >= 0 {

		if nums1[p1] > nums2[p2] {
			nums1[index] = nums1[p1]
			index -= 1
			p1 -= 1
		} else {
			nums1[index] = nums2[p2]
			index -= 1
			p2 -= 1
		}
	}
	for p2 >= 0 {
		nums1[index] = nums2[p2]
		p2 -= 1
		index -= 1
	}

}
