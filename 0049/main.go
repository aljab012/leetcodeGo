package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	wMap := make(map[string][]string)

	for _, word := range strs {
		sWord := SortString(word)
		wMap[sWord] = append(wMap[sWord], word)
	}
	for _, val := range wMap {
		ret = append(ret, val)
	}
	return ret
}

func SortString(s string) string {
	arr := strings.Split(s, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}
