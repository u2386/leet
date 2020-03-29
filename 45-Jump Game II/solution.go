package main

import (
	"fmt"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func jump(nums []int) int {
	var step, farthest, end int
	for i, n := range nums[:len(nums)-1] {
		farthest = max(farthest, i+n)
		if i == end {
			end = farthest
			step++
		}
	}
	return step
}

func main() {
	arr := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(arr))
}
