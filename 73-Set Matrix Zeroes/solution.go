package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	var row, col bool
	for i := 0; i < len(matrix); i++ {
		row = row || matrix[i][0] == 0
	}
	for i := 0; i < len(matrix[0]); i++ {
		col = col || matrix[0][i] == 0
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if col {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}

func main() {
	arr := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	setZeroes(arr)
	fmt.Println(arr)
}
