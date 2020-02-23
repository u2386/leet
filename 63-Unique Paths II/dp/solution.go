package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	obstacleGrid[0][0] = 1
	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			obstacleGrid[0][i] = 0
		} else {
			obstacleGrid[0][i] = obstacleGrid[0][i-1]
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}

	rows := len(obstacleGrid)
	cols := len(obstacleGrid[rows-1])
	return obstacleGrid[rows-1][cols-1]
}

func main() {
	grid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(grid))
}
