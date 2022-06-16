package main

import "strconv"
import "fmt"

// 8 9 9
// 8 7 9
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	arr1 := []rune(num1)
	arr2 := []rune(num2)

	Reverse(arr1, 0, len(arr1)-1)
	Reverse(arr2, 0, len(arr2)-1)

	length := len(arr1) + len(arr2)
	ret := make([]int, length)
	out := ""
	for i := range arr1 {
		for j := range arr2 {
			digit := int((arr1[i] - '0') * (arr2[j] - '0'))
			ret[i+j] += digit
			ret[i+j+1] += ret[i+j] / 10
			ret[i+j] = ret[i+j] % 10
		}
	}
	for i := len(ret) - 1; i >= 0; i-- {
		if ret[i] != 0 {
			ret = ret[:i+1]
			break
		}
	}
	for _, v := range ret {
		out = strconv.Itoa(v) + out
	}
	return out
}

func Swap(arr []rune, x, y int) []rune {
	if len(arr) == 1 {
		return arr
	}
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
	return arr
}

func Reverse(arr []rune, x, y int) []rune {
	for x < y {
		Swap(arr, x, y)
		x++
		y--
	}
	return arr
}

func main() {
	fmt.Println(multiply("123", "456"))
}
