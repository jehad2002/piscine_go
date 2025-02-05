package main

import (
	"fmt"
)

const N = 9 // Sudoku grid size

// Function to check if it's safe to place a number 'num' at board[row][col]
func isSafe(board [N][N]int, row, col, num int) bool {
	// Check the current row and column
	for i := 0; i < N; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Check the 3x3 subgrid
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

// Function to solve the Sudoku puzzle using backtracking
func solveSudoku(board [N][N]int) bool {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= N; num++ {
					if isSafe(board, row, col, num) {
						board[row][col] = num

						if solveSudoku(board) {
							return true
						}

						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// Function to print the solved Sudoku grid
func printSudoku(board [N][N]int) {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			fmt.Printf("%d ", board[row][col])
		}
		fmt.Println()
	}
}

func main() {
	board := [N][N]int{
		{3, 9, 6, 2, 4, 5, 7, 8, 1},
		{1, 7, 8, 3, 6, 9, 5, 2, 4},
		{5, 2, 4, 8, 1, 7, 3, 9, 6},
		{2, 8, 7, 9, 5, 1, 6, 4, 3},
		{9, 3, 1, 4, 8, 6, 2, 7, 5},
		{4, 6, 5, 7, 2, 3, 9, 1, 8},
		{7, 1, 2, 6, 3, 8, 4, 5, 9},
		{6, 5, 9, 1, 7, 4, 8, 3, 2},
		{8, 4, 3, 5, 9, 2, 1, 6, 7},
	}

	if solveSudoku(board) {
		fmt.Println("Solved Sudoku:")
		printSudoku(board)
	} else {
		fmt.Println("No solution exists.")
	}
}
