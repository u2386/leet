package main

import (
	"fmt"
	"strconv"
)

var opts = [...]byte{'+', '-', '*'}

func isOpt(c byte) bool {
	for _, o := range opts {
		if o == c {
			return true
		}
	}
	return false
}

func diffWaysToCompute(input string) []int {
	if n, err := strconv.Atoi(input); err == nil {
		return []int{n}
	}

	ret := []int{}
	for i := 0; i < len(input); i++ {
		c := input[i]
		if !isOpt(c) {
			continue
		}

		left := diffWaysToCompute(input[:i])
		right := diffWaysToCompute(input[i+1:])

		for _, l := range left {
			for _, r := range right {
				if c == '+' {
					ret = append(ret, l+r)
				} else if c == '-' {
					ret = append(ret, l-r)
				} else if c == '*' {
					ret = append(ret, l*r)
				}
			}
		}
	}
	return ret
}

func main() {
	input := "2*3-4*5"
	fmt.Println(diffWaysToCompute(input))
}
