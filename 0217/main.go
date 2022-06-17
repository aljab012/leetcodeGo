package main

func containsDuplicate(nums []int) bool {
	nMap := make(map[int]bool)

	for _, n := range nums {
		_, found := nMap[n]
		if found {
			return true
		}
		nMap[n] = true
	}
	return false
}

func main() {

}
