package main

/*
 * using a map to store the number and check if the number is already in the map
 *  * Space complexity: O(n)
 *  * Time complexity: O(n)
 */
func topKFrequent(nums []int, k int) []int {
	freqMap := map[int]int{}
	for _, n := range nums {
		freqMap[n]++
	}

	ret := []int{}
	maxK := len(nums)
	for len(ret) < k {
		for n, freq := range freqMap {
			if freq == maxK {
				ret = append(ret, n)
			}
		}
		maxK--
	}

	return ret
}
