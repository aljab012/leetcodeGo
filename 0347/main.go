package main

func topKFrequent(nums []int, k int) []int {
	countMap := map[int]int{}
	for _, n := range nums {
		countMap[n]++
	}
	freqMap := map[int][]int{}
	for k, v := range countMap {
		freqMap[v] = append(freqMap[v], k)
	}

	ret := []int{}
	maxFreq := len(nums)
	for len(ret) < k && maxFreq > 0 {
		for _, v := range freqMap[maxFreq] {
			if len(ret) < k {
				ret = append(ret, v)
			}
		}
		maxFreq--
	}
	return ret
}
