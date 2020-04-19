package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	m := [][]int{}
	if n == 0 {
		return m
	}

	dirs := [4][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{0, -1},
		[2]int{-1, 0},
	}

	for i := 0; i < n; i++ {
		m = append(m, make([]int, n))
	}

	i, j, v, s := 0, -1, 0, 0
	for {
		d := dirs[s]
		i += d[0]
		j += d[1]

		v++
		m[i][j] = v
		if v == n*n {
			break
		}

		if i+d[0] >= n || i+d[0] < 0 || j+d[1] >= n || j+d[1] < 0 || m[i+d[0]][j+d[1]] != 0 {
			s = (s + 1) % 4
		}
	}
	return m
}

func main() {
	fmt.Println(generateMatrix(3))
}
