package main

/*
 * using a map to store the number and check if the number is already in the map
 *  * Space complexity: O(n)
 *  * Time complexity: O(n)
 */
func topKFrequent(nums []int, k int) []int {
	nToFreq := map[int]int{}
	for _, n := range nums {
		nToFreq[n]++
	}
	freqToN := map[int][]int{}
	for n, freq := range nToFreq {
		freqToN[freq] = append(freqToN[freq], n)
	}
	ret := []int{}
	for i := len(nums); i >= 0; i-- {
		for _, n := range freqToN[i] {
			if len(ret) < k {
				ret = append(ret, n)
			}
		}
	}
	return ret
}
