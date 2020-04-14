package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}

	i, j, mid := 0, len(nums)-1, 0
	for i <= j {
		mid = (i + j) >> 1
		if target == nums[mid] && (mid == 0 || nums[mid-1] < target) {
			ret[0] = mid
			break
		} else if target <= nums[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	i, j, mid = 0, len(nums)-1, 0
	for i <= j {
		mid = (i + j) >> 1
		if target == nums[mid] && (mid == len(nums)-1 || target < nums[mid+1]) {
			ret[1] = mid
			break
		} else if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return ret
}

func main() {
	arr := []int{2, 2}
	target := 8
	fmt.Println(searchRange(arr, target))
}
