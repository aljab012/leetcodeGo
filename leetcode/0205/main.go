package main

func isIsomorphic(s string, t string) bool {
	mapST := map[string]string{}
	mapTS := map[string]string{}
	for i := range s {
		c1 := string(s[i])
		c2 := string(t[i])
		if val, ok := mapST[string(c1)]; ok && string(c2) != val {
			return false
		}
		if val, ok := mapTS[string(c2)]; ok && string(c1) != val {
			return false
		}
		mapST[string(c1)] = string(c2)
		mapTS[string(c2)] = string(c1)
	}
	return true
}
