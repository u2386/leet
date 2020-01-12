package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	dirs := [][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}

	row, col := len(matrix), len(matrix[0])
	footprint := make([][]bool, row)
	for i := range matrix {
		footprint[i] = make([]bool, col)
	}

	path := []int{}
	x, y, d := 0, 0, 0
	dx, dy := dirs[d][0], dirs[d][1]
	for {
		if x < 0 || x > (row-1) || y < 0 || y > (col-1) || footprint[x][y] {
			break
		}

		footprint[x][y] = true
		path = append(path, matrix[x][y])

		if !(0 <= x+dx && x+dx <= (row-1)) || !(0 <= y+dy && y+dy <= (col-1)) || footprint[x+dx][y+dy] {
			d = (d + 1) % 4
			dx, dy = dirs[d][0], dirs[d][1]
		}
		x += dx
		y += dy
	}
	return path
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println(spiralOrder(matrix))
}
