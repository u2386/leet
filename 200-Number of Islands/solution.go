package main

import (
	"fmt"
)

type unionFindSet struct {
	count   int
	parents []int
	ranks   []int
}

func createUnionFindSet(n int) *unionFindSet {
	us := &unionFindSet{
		count:   n,
		parents: make([]int, n),
		ranks:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		us.parents[i] = i
	}
	return us
}

func (us *unionFindSet) union(x, y int) {
	xp, yp := us.find(x), us.find(y)
	if xp == yp {
		return
	}
	if us.ranks[xp] < us.ranks[yp] {
		us.parents[xp] = yp
	} else if us.ranks[xp] > us.ranks[yp] {
		us.parents[yp] = xp
	} else {
		us.parents[yp] = xp
		us.ranks[xp]++
	}
	us.count--
}

func (us *unionFindSet) find(x int) int {
	if x != us.parents[x] {
		us.parents[x] = us.find(us.parents[x])
	}
	return us.parents[x]
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	index := func(x, y int) int {
		return x*col + y
	}

	dirs := [][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
	}
	dummy := row * col
	uf := createUnionFindSet(row*col + 1)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				uf.union(index(i, j), dummy)
			} else {
				for _, d := range dirs {
					di, dj := i+d[0], j+d[1]
					if di < row && dj < col && grid[di][dj] == '1' {
						uf.union(index(di, dj), index(i, j))
					}
				}
			}
		}
	}

	return uf.count - 1
}

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
