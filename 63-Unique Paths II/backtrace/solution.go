package main

import (
	"fmt"
)

func findPaths(grid [][]int, x, y int, count *int) {
	if x == len(grid)-1 && y == len(grid[x])-1 {
		*count++
		return
	}
	if x+1 < len(grid) && grid[x+1][y] == 0 {
		findPaths(grid, x+1, y, count)
	}
	if y+1 < len(grid[x]) && grid[x][y+1] == 0 {
		findPaths(grid, x, y+1, count)
	}
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	count := 0
    if obstacleGrid[0][0] == 1 {
        return 0
    }
	findPaths(obstacleGrid, 0, 0, &count)
	return count
}

func main() {
	grid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(grid))
}
