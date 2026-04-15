package main

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(0, &board)
}

// Try to place a queen in column col
func solve(col int, board *[8]int) {
	if col == 8 {
		printBoard(board)
		return
	}

	for row := 1; row <= 8; row++ {
		if safe(col, row, board) {
			board[col] = row
			solve(col+1, board)
		}
	}
}

// Check if placing a queen at (col, row) is safe
func safe(col, row int, board *[8]int) bool {
	for c := 0; c < col; c++ {
		r := board[c]
		if r == row || // same row
			r-(c) == row-col || // descending diagonal
			r+(c) == row+col { // ascending diagonal
			return false
		}
	}
	return true
}

// Print the solution as 8 digits
func printBoard(board *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '0'))
	}
	z01.PrintRune('\n')
}

func main2() {
	EightQueens()
}