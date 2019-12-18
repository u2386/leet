package main

import "fmt"

func pourWater(heights []int, V, K int) []int {
	for i := 0; i < V; i++ {
		k := K

		for j := k - 1; j >= 0; j-- {
			if heights[j] > heights[k] {
				break
			}
			if heights[j] < heights[k] {
				k = j
			}
		}
		if k != K {
			heights[k]++
			continue
		}

		for j := k + 1; j < len(heights); j++ {
			if heights[j] > heights[k] {
				break
			}
			if heights[j] < heights[k] {
				k = j
				break
			}
		}
		heights[k]++
	}
	return heights
}

func main() {
	heights := []int{2, 1, 1, 2, 1, 2, 2}
	fmt.Println(pourWater(heights, 4, 3))
}
