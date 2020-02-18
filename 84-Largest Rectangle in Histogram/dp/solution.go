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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	maxRect := 0
	var lowest int
	for i := 0; i < len(heights); i++ {
		if heights[i] == 0 {
			continue
		}
		lowest = heights[i]
		for j := i; j < len(heights); j++ {
			if heights[j] == 0 {
				break
			}
			lowest = min(heights[j], lowest)

			if j == i {
				maxRect = max(heights[j], maxRect)
			} else {
				maxRect = max(lowest*(j-i+1), maxRect)
			}
		}
	}
	return maxRect
}

func main() {
	arr := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea((arr)))
}
