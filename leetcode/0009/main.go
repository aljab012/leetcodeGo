package main

func isPalindrome(x int) bool {
	reversed := 0
	for num := x; num > 0; num = num / 10 {
		reversed = reversed*10 + num%10
	}
	return reversed == x
}
