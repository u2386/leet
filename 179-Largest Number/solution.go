package main

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { 
	x := strconv.Itoa(a[i])
	y := strconv.Itoa(a[j])
	if len(x) < len(y) {
		return false
	} else if len(x) > len(y) {
		return true
	}
	return x < y
 }

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	} else if len(nums) == 1 {
		return string(nums[0])
	}

	sort.Sort(sort.Reverse(SortBy(nums)))
	var sb strings.Builder
	for _, num := range nums {
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
}

func main() {
	fmt.Println(largestNumber([]int{3,30,34,5,9}))
}
