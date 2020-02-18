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

func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	var h int
	maxRect := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			// 当前最大高度
			h, stack = heights[stack[len(stack)-1]], stack[:len(stack)-1]
			// 在最大高度下的最大宽度
			// 如果 stack 不为空，因为 stack 递增，矩形宽左侧只到前一个元素位置；
			// 如果 stack 为空，矩形宽左侧到最近一个小于当前元素位置
			maxRect = max(maxRect, h*(i-stack[len(stack)-1]-1))
		}
		// 保持递增
		stack = append(stack, i)
	}

	for len(stack) > 1 {
		h, stack = heights[stack[len(stack)-1]], stack[:len(stack)-1]
		maxRect = max(maxRect, h*(len(heights)-stack[len(stack)-1]-1))
	}

	return maxRect
}

func main() {
	arr := []int{1, 3, 2, 3, 1}
	fmt.Println(largestRectangleArea((arr)))
}
