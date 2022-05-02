import ("fmt")

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits

}

func main(){
	arr := [9]
	fmt.Println(plusOne(arr))
}
