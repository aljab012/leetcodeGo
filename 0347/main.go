func topKFrequent(nums []int, k int) []int {
	nToCount := map[int]int{}
	countToN := map[int][]int{}

	for _, n := range nums {
		nToCount[n] += 1
	}
	for n, count := range nToCount {
		countToN[count] = append(countToN[count], n)
	}
	ret := []int{}

	for count := len(nums); len(ret) != k; count-- {
		for _, n := range countToN[count] {
			if len(ret) < k {
				ret = append(ret, n)

			}
		}
	}
	return ret
}
