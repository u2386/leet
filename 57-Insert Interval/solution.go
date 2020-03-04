package main

import (
	"fmt"
	"sort"
)

// SortBy ...
type SortBy [][]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i][0] < a[j][0] }

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)

	if len(intervals) == 1 {
		return intervals
	}

	sort.Sort(SortBy(intervals))
	fmt.Println(intervals)

	n := 0
	for i, p := range intervals[1:] {
		if intervals[n][1] < p[0] {
			n++
			intervals[n] = intervals[i+1]
			continue
		}

		intervals[n][0] = min(intervals[n][0], p[0])
		intervals[n][1] = max(intervals[n][1], p[1])
	}
	return intervals[:n+1]
}

func main() {
	intervals := [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}
	newInterval := []int{4, 8}
	fmt.Println(insert(intervals, newInterval))
}
