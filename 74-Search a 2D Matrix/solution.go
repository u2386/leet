package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	
	rows, cols := len(matrix), len(matrix[0])
	i, j, mid := 0, rows*cols-1, 0

	for i <= j {
		mid = (i + j) >> 1
		x, y := mid/cols, mid%cols
		v := matrix[x][y]
		if v == target {
			return true
		} else if v < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return false
}

func main() {
	arr := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	target := 11
	fmt.Println(searchMatrix(arr, target))
}
