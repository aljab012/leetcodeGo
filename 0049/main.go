package main

func groupAnagrams(strs []string) [][]string {
	groupMap := map[[26]int][]string{}
	for _, s := range strs {
		arr := [26]int{}
		for _, c := range s {
			arr[c-'a']++
		}
		groupMap[arr] = append(groupMap[arr], s)
	}

	ret := [][]string{}
	for _, v := range groupMap {
		ret = append(ret, v)
	}
	return ret
}
