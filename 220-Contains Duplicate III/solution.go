package main

import (
	"fmt"
)

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 || k < 0 {
		return false
	}

	bucket := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := nums[i] / (t + 1)
		if v, ok := bucket[n]; ok && abs(v-nums[i]) <= t {
			return true
		}
		if v, ok := bucket[n-1]; ok && abs(v-nums[i]) <= t {
			return true
		}
		if v, ok := bucket[n+1]; ok && abs(v-nums[i]) <= t {
			return true
		}
		bucket[n] = nums[i]
		if i >= k {
			delete(bucket, nums[i-k]/(t+1))
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 1}
	k := 3
	t := 0
	fmt.Println(containsNearbyAlmostDuplicate(arr, k, t))
}
