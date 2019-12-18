package main

import "fmt"

func pourWater(heights []int, V, K int) []int {
	var j int
	for i := 0; i < V; i++ {
		k := K

		j = k
		for j > 0 && heights[j] >= heights[j-1] {
			j--
			if heights[j] < heights[k] {
				k = j
			}
		}

		j = k
		for j < len(heights)-1 && heights[j] >= heights[j+1] {
			j++
			if heights[j] < heights[k] {
				k = j
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
