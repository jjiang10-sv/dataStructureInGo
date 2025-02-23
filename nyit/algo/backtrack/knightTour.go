package main

import "fmt"

const N = 8

// Possible moves for a knight
var knightMoves = [][2]int{
	{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
}

// Print the chessboard
func printBoardKnight(board [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%2d ", board[i][j])
		}
		fmt.Println()
	}
}

// Check if the move is within bounds and the square is unvisited
func isSafeKnight(x, y int, board [N][N]int) bool {
	return x >= 0 && y >= 0 && x < N && y < N && board[x][y] == -1
}

// Solve the Knight's Tour using backtracking
func knightTour(board *[N][N]int, x, y, move int) bool {
	// If all squares are visited, return true
	if move == N*N {
		return true
	}

	// Try all possible moves
	for _, m := range knightMoves {
		nextX := x + m[0]
		nextY := y + m[1]

		if isSafeKnight(nextX, nextY, *board) {
			board[nextX][nextY] = move
			if knightTour(board, nextX, nextY, move+1) {
				return true
			}
			// Backtrack
			board[nextX][nextY] = -1
		}
	}

	return false
}

func mainKnight() {
	var board [N][N]int
	// Initialize the board
	for i := range board {
		for j := range board[i] {
			board[i][j] = -1
		}
	}

	// Starting position
	startX, startY := 0, 0
	board[startX][startY] = 0

	// Solve the tour
	if knightTour(&board, startX, startY, 1) {
		fmt.Println("Knight's Tour solution:")
		printBoardKnight(board)
	} else {
		fmt.Println("No solution exists.")
	}
}