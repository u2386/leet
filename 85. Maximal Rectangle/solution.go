package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	var ret int

	H, W := len(matrix), len(matrix[0])
	heights := make([]int, W)

	L, R := make([]int, W), make([]int, W)
	for i := 0; i < W; i++ {
		L[i] = -1
		R[i] = W
	}

	boundary := -1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		boundary = -1
		for j := 0; j < W; j++ {
			if heights[j] > 0 {
				L[j] = max(L[j], boundary)
			} else {
				L[j] = -1
				boundary = j
			}
		}

		boundary = W
		for j := W-1; j >= 0; j-- {
			if heights[j] > 0 {
				R[j] = min(R[j], boundary)
			} else {
				R[j] = W
				boundary = j
			}
		}

		area := 0
		for j := 0; j < W; j++ {
			area = (R[j] - L[j] - 1) * heights[j]
			ret = max(area, ret)
		}
	}
	return ret
}

func main() {
	m := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}

	fmt.Println(maximalRectangle(m))
}
