package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type SortBy []int

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	var x, y strings.Builder

	x.Reset()
	y.Reset()

	x.WriteString(strconv.Itoa(a[i]))
	x.WriteString(strconv.Itoa(a[j]))

	y.WriteString(strconv.Itoa(a[j]))
	y.WriteString(strconv.Itoa(a[i]))

	return x.String() < y.String()
}

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	} else if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}

	flag := true
	sort.Sort(sort.Reverse(SortBy(nums)))
	var sb strings.Builder
	for _, num := range nums {
		if flag && num == 0 {
			continue
		}
		flag = false
		sb.WriteString(strconv.Itoa(num))
	}
	ret := sb.String()
	if len(ret) == 0 {
		return "0"
	}
	return ret
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
