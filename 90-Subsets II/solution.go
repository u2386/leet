package main

import (
	"fmt"
)

type retval struct {
	accu  [][]int
	pivot int
}

func subsets(nums []int) retval {
	ret := retval{
		accu:  [][]int{},
		pivot: 0,
	}

	if len(nums) == 0 {
		ret.accu = append(ret.accu, nums)
		ret.pivot++
		return ret
	}

	dup := len(nums) > 1 && nums[0] == nums[1]
	prev := subsets(nums[1:])
	for _, v := range prev.accu {
		ret.accu = append(ret.accu, v)
	}
	for i, v := range prev.accu {
		if dup && i < prev.pivot {
			continue
		}
		ret.accu = append(ret.accu, append([]int{nums[0]}, v...))
	}
	ret.pivot = len(prev.accu)
	return ret
}

func subsetsWithDup(nums []int) [][]int {
	prev := subsets(nums)
	return prev.accu
}

func main() {
	arr := []int{1, 1, 2, 2}
	fmt.Println(subsetsWithDup(arr))
}
