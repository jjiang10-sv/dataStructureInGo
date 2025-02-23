package main

import (
	"fmt"
	"math"
)
// Explanation

// 	1.	Representation:
// 	•	The board is represented as a one-dimensional array, where the index is the row, and the value is the column where the queen is placed.
// 	2.	Checking Safety:
// 	•	The isSafe function checks:
// 	•	If another queen exists in the same column.
// 	•	If another queen exists on the same diagonal (both directions).
// 	3.	Recursive Backtracking:
// 	•	The solveNQueens function tries placing queens row by row.
// 	•	If placing a queen in a column leads to a valid configuration, the function proceeds to the next row.
// 	•	If not, it backtracks and tries a different column.
// 	4.	Base Case:
// 	•	When all rows are filled, a valid solution is found and stored.
// 	5.	Printing Solutions:
// 	•	The solutions are stored in a slice of slices and printed in a human-readable chessboard format.

// Complexity

// 	•	Time Complexity:
// 	•	O(N!): For an N \times N board, there are N! possible placements in the worst case.
// 	•	The backtracking reduces this by pruning invalid configurations.
// 	•	Space Complexity:
// 	•	O(N) for the board array and recursive stack.

// This code generates all possible solutions to the 8-Queens Puzzle and prints them, illustrating the elegance of the backtracking algorithm.
func printBoard(board []int) {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i] == j {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func isSafe1(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || 
			math.Abs(float64(board[i]-col)) == math.Abs(float64(i-row)) {
			return false
		}
	}
	return true
}

func solveNQueens(board []int, row int, solutions *[][]int) {
	n := len(board)
	if row == n {
		// Store the solution
		solution := make([]int, n)
		copy(solution, board)
		*solutions = append(*solutions, solution)
		return
	}

	for col := 0; col < n; col++ {
		if isSafe1(board, row, col) {
			board[row] = col
			solveNQueens(board, row+1, solutions)
			board[row] = -1 // Backtrack
		}
	}
}

func main() {
	n := 8 // Size of the board
	board := make([]int, n)
	for i := range board {
		board[i] = -1
	}

	var solutions [][]int
	solveNQueens(board, 0, &solutions)

	fmt.Printf("Found %d solutions:\n", len(solutions))
	for _, solution := range solutions {
		printBoard(solution)
	}
}