package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i-1 >= 0 && j-1 >= 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			} else if i-1 >= 0 {
				grid[i][j] += grid[i-1][j]
			} else if j-1 >= 0 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {

	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
