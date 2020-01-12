package main

import (
	"fmt"
)

type pos struct {
	x int
	y int
}

func explore(p pos, grid [][]byte, seen [][]bool) {
	q := []pos{}
	q = append(q, p)
	for len(q) != 0 {
		p, q = q[0], q[1:]
		if grid[p.x][p.y] == '1' {
			if p.x-1 >= 0 && !seen[p.x-1][p.y] {
				q = append(q, pos{p.x - 1, p.y})
				seen[p.x-1][p.y] = true
			}
			if p.x+1 < len(grid) && !seen[p.x+1][p.y] {
				q = append(q, pos{p.x + 1, p.y})
				seen[p.x+1][p.y] = true

			}
			if p.y-1 >= 0 && !seen[p.x][p.y-1] {
				q = append(q, pos{p.x, p.y - 1})
				seen[p.x][p.y-1] = true

			}
			if p.y+1 < len(grid[p.x]) && !seen[p.x][p.y+1] {
				q = append(q, pos{p.x, p.y + 1})
				seen[p.x][p.y+1] = true
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	seen := make([][]bool, len(grid))
	for i := range grid {
		seen[i] = make([]bool, len(grid[i]))
	}

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !seen[i][j] {
				seen[i][j] = true
				explore(pos{i, j}, grid, seen)
				count++
			}
		}
	}
	return count
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
