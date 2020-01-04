package main

import "fmt"

func bsearch(nums []int, start, end int, target int) int {
	for start < end {
		mid := (start + end) >> 1
		if target <= nums[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tail := []int{nums[0]}
	for _, n := range nums[1:] {
		pos := bsearch(tail, 0, len(tail), n)

		if pos == len(tail) {
			tail = append(tail, n)
		} else {
			tail[pos] = n
		}
	}
	return len(tail)
}

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(arr))
}
