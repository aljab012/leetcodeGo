func singleNonDuplicate(nums []int) int {
    ret := 0

    for _, n := range nums{
        ret = ret ^ n
    }
    return ret
}