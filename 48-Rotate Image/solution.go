package main

import "fmt"

func rotate(matrix [][]int) {
	N := len(matrix)
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N>>1; j++ {
			matrix[i][j], matrix[i][N-j-1] = matrix[i][N-j-1], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
