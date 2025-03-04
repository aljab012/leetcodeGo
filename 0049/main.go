package main

import "sort"

// with sorting
func groupAnagrams1(strs []string) [][]string {
	freqMap := map[string][]string{}

	for _, str := range strs {
		sorted := sortString(str)
		freqMap[sorted] = append(freqMap[sorted], str)
	}

	ret := [][]string{}
	for _, v := range freqMap {
		ret = append(ret, v)
	}
	return ret
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// with frequency array
func groupAnagrams2(strs []string) [][]string {
	freqMap := map[[26]int][]string{}

	for _, str := range strs {
		chatCount := [26]int{}
		for _, c := range str {
			chatCount[c-'a']++
		}
		freqMap[chatCount] = append(freqMap[chatCount], str)
	}

	ret := [][]string{}
	for _, v := range freqMap {
		ret = append(ret, v)
	}
	return ret
}
