package main

import "fmt"

func permute(nums []int) [][]int {
	ret := [][]int{}

	if len(nums) == 0 {
		return ret
	} else if len(nums) == 1 {
		ret = append(ret, []int{nums[0]})
		return ret
	}

	seen := make([]bool, len(nums))
	ret = _permute(nums, ret, []int{}, seen)
	return ret
}

func _permute(nums []int, ret [][]int, path []int, seen []bool) [][]int {
	if len(path) == len(nums) {
		ret = append(ret, append(path[:0:0], path...))
		return ret
	}

	for i, n := range nums {
		if v := seen[i]; v {
			continue
		}
		seen[i] = true
		path = append(path, n)
		ret = _permute(nums, ret, path, seen)
		path = path[:len(path)-1]
		seen[i] = false
	}
	return ret
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(permute(arr))
}
