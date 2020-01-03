package main

import (
	"fmt"
	"sort"
)

type sortBy []int

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[i] < a[j] }

func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(sortBy(candidates))
	return combinationSum(candidates, 0, target, []int{}, [][]int{})
}

func combinationSum(candidates []int, start int, target int, p []int, ret [][]int) [][]int {
	if target == 0 {
		ret = append(ret, append(p[:0:0], p...))
		return ret
	}

	if start == len(candidates) {
		return ret
	}

	for i := start; i < len(candidates); i++ {
		n := candidates[i]
		if target-n < 0 {
			break
		}
		if i > start && n == candidates[i-1] {
			continue
		}
		p = append(p, n)
		ret = combinationSum(candidates, i+1, target-n, p, ret)
		p = p[:len(p)-1]
	}
	return ret
}

func main() {
	arr := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(arr, 8))
}
