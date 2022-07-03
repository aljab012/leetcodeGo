package main

import "fmt"
import "strconv"

func evalRPN(tokens []string) int {

	funcMap := map[string]func(x, y int) int{
		"*": func(x, y int) int { return x * y },
		"/": func(x, y int) int { return x / y },
		"+": func(x, y int) int { return x + y },
		"-": func(x, y int) int { return x - y },
	}

	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			l := stack[len(stack)-2]
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, funcMap[token](l, r))
		}
	}
	return stack[0]
}

// ["2","1","+","3","*"]
//  ((2 + 1) * 3) = 9
func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}
