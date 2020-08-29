package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestValidParentheses(s string) int {
	stack := []int{-1}
	ret := 0
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				ret = max(ret, i - top)
			} else {
				stack = append(stack, i)
			}
		}
	}
	return ret
}

func main()  {
	fmt.Println(longestValidParentheses(")()())"))
}