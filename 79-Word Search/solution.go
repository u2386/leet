package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	var dirs = [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}

	var find func(board [][]byte, word string, px, py int) bool
	find = func(board [][]byte, word string, px, py int) bool {
		if board[px][py] != word[0] {
			return false
		}

		if len(word) == 1 {
			return true
		}

		var orig byte
		orig, board[px][py] = board[px][py], orig
		for i := 0; i < len(dirs); i++ {
			dx := dirs[i][0]
			dy := dirs[i][1]

			nx, ny := px+dx, py+dy
			if nx >= 0 && nx < len(board) && ny >= 0 && ny < len(board[nx]) {
				if find(board, word[1:], nx, ny) {
					return true
				}
			}
		}
		board[px][py] = orig
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if find(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		[]byte{'A', 'A'},
	}

	word := "AAA"
	fmt.Println(exist(board, word))
}
