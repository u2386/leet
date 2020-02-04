package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	digits[len(digits)-1]++

	if digits[len(digits)-1] > 9 {
		carry := 0
		for i := len(digits) - 1; i >= 0; i-- {
			digits[i] += carry
			carry = digits[i] / 10
			if carry == 0 {
				break
			}
			digits[i] = digits[i] % 10
		}

		if carry > 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}

func main() {
	arr := []int{4, 3, 2, 1}
	fmt.Println(plusOne(arr))
}
