package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	sets := subsets(nums[1:])
	accu := [][]int{}
	for _, s := range sets {
		accu = append(accu, s)
		accu = append(accu, append([]int{nums[0]}, s...))
	}
	return append([][]int{}, accu...)
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(subsets(arr))
}
