package main

import (
	"fmt"
)

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	stack := []int{}
	stack = append(stack, 0)
	stack = append(stack, len(nums)-1)

	for len(stack) > 0 {
		high := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		low := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if low < 0 || high >= len(nums) || !(low < high) {
			continue
		}

		i, j, pivot := low, high, nums[low]
		for i < j {
			for i < j && pivot <= nums[j] {
				j--
			}
			nums[i] = nums[j]

			for i < j && nums[i] <= pivot {
				i++
			}
			nums[j] = nums[i]
		}

		nums[i] = pivot

		stack = append(stack, low)
		stack = append(stack, i-1)
		stack = append(stack, j+1)
		stack = append(stack, high)
	}
}

func main() {
	arr := []int{39, 28, 55, 87, 66, 3, 17, 39}
	qsort(arr)
	fmt.Println(arr)
}
