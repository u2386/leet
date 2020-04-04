package main

import (
	"sort"
	"fmt"
)

var prev [][]int

func subsets(nums []int) [][]int {
	accu := [][]int{}
	if len(nums) == 0 {
		return append(accu, []int{})
	}

	sets := subsets(nums[1:])
	for _, v := range sets {
		accu = append(accu, v)
	}

	if !(len(nums) > 1 && nums[0] == nums[1]) {
		prev = sets
	}
	p := append(prev[:0:0], prev...)
	prev = [][]int{}
	for _, v := range p {
		v := append([]int{nums[0]}, v...)
		accu = append(accu, v)
		prev = append(prev, v)
	}
	return accu
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return subsets(nums)
}

func main() {
	arr := []int{4, 4, 4, 1, 4}
	fmt.Println(subsetsWithDup(arr))
}
