package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
		if len(m) > k {
			delete(m, nums[i-k])
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(arr, k))
}
