package main

import (
	"fmt"
)

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	start, end, i := 0, len(nums)-1, 0
	for start <= i && i <= end {
		if nums[i] == 0 {
			nums[i] = nums[start]
			nums[start] = 0
			start++
			i++
		} else if nums[i] == 2 {
			nums[i] = nums[end]
			nums[end] = 2
			end--
		} else {
			i++
		}
	}
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
}
