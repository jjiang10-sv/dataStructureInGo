package main

import (
	"fmt"
)

// Size of the Sudoku grid
const SIZE = 9

// Sudoku grid type definition
type Sudoku [SIZE][SIZE]int

// printGridSoduko prints the Sudoku grid
func printGridSoduko(grid Sudoku) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
}

// isSafeSo checks if it's safe to place a number at a given position
func isSafeSo(grid Sudoku, row, col, num int) bool {
	// Check row
	for i := 0; i < SIZE; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	// Check column
	for i := 0; i < SIZE; i++ {
		if grid[i][col] == num {
			return false
		}
	}

	// Check 3x3 subgrid
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

// solveSudoku uses backtracking to solve the Sudoku puzzle
func solveSudoku(grid *Sudoku) bool {
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
			if grid[row][col] == 0 { // Find empty cell
				for num := 1; num <= SIZE; num++ { // Try numbers 1-9
					if isSafeSo(*grid, row, col, num) {
						grid[row][col] = num
						if solveSudoku(grid) {
							return true
						}
						// If no solution, reset and backtrack
						grid[row][col] = 0
					}
				}
				return false // If no valid number found
			}
		}
	}
	return true // Puzzle solved
}

func mainSo() {
	// Example Sudoku puzzle (0 represents empty cells)
	grid := Sudoku{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	fmt.Println("Initial Sudoku Puzzle:")
	printGridSoduko(grid)

	if solveSudoku(&grid) {
		fmt.Println("\nSolved Sudoku Puzzle:")
		printGridSoduko(grid)
	} else {
		fmt.Println("\nNo solution exists!")
	}
}

// Explanation

// 	1.	Grid Representation:
// 	•	The Sudoku grid is represented as a 2D array Sudoku[9][9] where each element can be a number between 1 to 9, or 0 for empty cells.
// 	2.	isSafe function:
// 	•	This function ensures that placing a number at a given cell is safe by checking if the number already exists in the current row, column, or 3x3 subgrid.
// 	3.	Backtracking:
// 	•	The solveSudoku function iterates through each cell, attempting to place a valid number (1-9) using the isSafe function.
// 	•	If the number is valid, it places it in the cell and recursively attempts to solve the puzzle.
// 	•	If a conflict arises, the algorithm resets the cell and backtracks.
// 	4.	Output:
// 	•	The printGrid function is used to display the grid before and after solving.

// Key Concepts

// 	•	Backtracking: The backtracking algorithm explores all possible combinations of numbers in the grid, attempting to solve the puzzle step by step. If a conflict is encountered, it backtracks to the previous state and tries a different number.
// 	•	Efficiency: This approach is guaranteed to find a solution if one exists, but may not be the most efficient for larger or more complex puzzles. Advanced techniques like constraint propagation can improve performance, but backtracking remains a fundamental method for solving constraint satisfaction problems like Sudoku.
