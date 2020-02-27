package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	for n&15 == 0 {
		n >>= 4
	}
	for _, i := range []int{0, 1, 2, 4, 8} {
		if n == i {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPowerOfTwo(16))
}
