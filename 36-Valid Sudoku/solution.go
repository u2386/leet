package main

import (
	"fmt"
	"strconv"
)

func checkRows(board [][]byte) bool {
	var seen []bool
	var i int
	for _, row := range board {
		seen = make([]bool, 10)
		for _, v := range row {
			if v != '.' {
				i, _ = strconv.Atoi(string(v))
				if seen[i] {
					return false
				}
				seen[i] = true
			}
		}
	}
	return true
}

func checkCols(board [][]byte) bool {
	var seen []bool
	var v byte
	var i int
	for c := 0; c < len(board[0]); c++ {
		seen = make([]bool, 10)
		for r := 0; r < len(board); r++ {
			v = board[r][c]
			if v != '.' {
				i, _ = strconv.Atoi(string(v))
				if seen[i] {
					return false
				}
				seen[i] = true
			}
		}
	}
	return true
}

func checkGrids(board [][]byte) bool {
	var seen []bool
	var v byte
	var i int
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			seen = make([]bool, 10)
			for k := 0; k < 9; k++ {
				v = board[x*3+(k/3)][y*3+(k%3)]
				if v != '.' {
					i, _ = strconv.Atoi(string(v))
					if seen[i] {
						return false
					}
					seen[i] = true
				}
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	return checkRows(board) && checkCols(board) && checkGrids(board)
}

func main() {
	board := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board))
}
