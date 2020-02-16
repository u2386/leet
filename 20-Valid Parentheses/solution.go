package main

import "fmt"

var open = map[byte]int{
	'(': 0,
	'[': 1,
	'{': 2,
}

var close = map[byte]int{
	')': 0,
	']': 1,
	'}': 2,
}

func isValid(s string) bool {
	stack := []byte{}
	var c byte
	for len(s) > 0 {
		c, s = s[0], s[1:]

		if _, ok := open[c]; ok {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 || close[c] != open[stack[len(stack)-1]] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
