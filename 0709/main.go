package main

func toLowerCase(s string) string {
	diff := 'a' - 'A'
	arr := []rune(s)
	for i := range arr {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = arr[i] + diff
		}
	}
	return string(arr)
}
