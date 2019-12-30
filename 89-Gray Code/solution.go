package main

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	prev := grayCode(n - 1)
	h := 1 << uint(n-1)
	ret := append(prev[:0:0], prev...)

	for i := (len(prev) >> 1) - 1; i >= 0; i-- {
		opp := len(prev) - 1 - i
		prev[i], prev[opp] = prev[opp], prev[i]
	}
	for _, x := range prev {
		ret = append(ret, x|h)
	}
	return ret
}

func main() {
	fmt.Println(grayCode(2))
}
