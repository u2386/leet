package main

import (
	"fmt"
)

const (
	INT_MAX = 2147483647
	INT_MIN = -2147483648
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func myAtoi(str string) int {
	var ret int
	sign := 0
	hit := false
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			hit = true
			ret = ret*10 + int(str[i]-'0')
			if ret > 2147483648 {
				break
			}
		} else if !hit && sign == 0 && str[i] == '-' {
			hit = true
			sign = -1
		} else if !hit && sign == 0 && str[i] == '+' {
			hit = true
			sign = 1
		} else if hit {
			break
		} else if !hit && str[i] == ' ' {
			continue
		} else {
			return 0
		}
	}

	if sign != 0 {
		ret *= sign
	}
	return min(INT_MAX, max(INT_MIN, ret))
}

func main() {
	fmt.Println(int32(INT_MAX))
	fmt.Println(int32(INT_MIN))

	s := "42"
	fmt.Println(myAtoi(s))
}
