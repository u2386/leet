package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	i, j, mid := 0, len(nums)-1, 0
	for i <= j {
		mid = (i + j) >> 1
		if target < nums[mid] {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			return mid
		}
	}

	if target < nums[mid] {
		return mid
	}
	return mid + 1
}

func main() {
	arr := []int{1, 3, 5, 6}
	target := 4
	fmt.Println(searchInsert(arr, target))
}
