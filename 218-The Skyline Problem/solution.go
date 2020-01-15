package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSkyline(buildings [][]int) [][]int {
	skyline := [][]int{}

	if len(buildings) == 0 {
		return skyline
	} else if len(buildings) == 1 {
		b := buildings[0]
		skyline = append(skyline, []int{b[0], b[2]})
		skyline = append(skyline, []int{b[1], 0})
		return skyline
	}

	mid := len(buildings) >> 1
	left := getSkyline(buildings[:mid])
	right := getSkyline(buildings[mid:])

	var x []int
	lh, rh := 0, 0
	for len(left) > 0 && len(right) > 0 {
		l, r := left[0], right[0]

		if l[0] < r[0] {
			left = left[1:]
			x = []int{l[0], max(l[1], rh)}
			lh = l[1]
		} else if l[0] > r[0] {
			right = right[1:]
			x = []int{r[0], max(r[1], lh)}
			rh = r[1]
		} else {
			left = left[1:]
			right = right[1:]
			x = []int{l[0], max(l[1], r[1])}
			lh = l[1]
			rh = r[1]
		}

		if len(skyline) == 0 || skyline[len(skyline)-1][1] != x[1] {
			skyline = append(skyline, x)
		}
	}
	skyline = append(skyline, left...)
	skyline = append(skyline, right...)
	return skyline
}

func main() {
	buildings := [][]int{
		[]int{2, 9, 10},
		[]int{3, 7, 15},
		[]int{5, 12, 12},
		[]int{15, 20, 10},
		[]int{19, 24, 8},
	}
	fmt.Println(getSkyline(buildings))
}
