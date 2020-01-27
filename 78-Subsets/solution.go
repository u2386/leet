package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	N := 1 << uint(len(nums))
	ret := [][]int{}
	for i := 0; i < N; i++ {
		p := []int{}
		for j := 0; j < len(nums); j++ {
			if (i & (1 << uint(j))) != 0 {
				p = append(p, nums[j])
			}
		}
		ret = append(ret, p)
	}
	return ret
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
