package main

import "fmt"
// To solve a maze using backtracking in Go, you can use a depth-first search (DFS) approach. The algorithm explores potential paths in the maze and backtracks whenever it reaches a dead end. Here’s a step-by-step explanation and implementation.

// Backtracking Algorithm for Solving a Maze

// 	1.	Problem Setup:
// 	•	Represent the maze as a 2D grid where cells can be:
// 	•	0: Path (walkable)
// 	•	1: Wall (blocked)
// 	•	Define a starting point (e.g., top-left) and an ending point (e.g., bottom-right).
// 	2.	Algorithm:
// 	•	Begin at the starting cell.
// 	•	Mark the current cell as visited to avoid re-visiting.
// 	•	Recursively attempt to move in all four directions (up, down, left, right).
// 	•	If the move reaches the ending point, return success.
// 	•	If no moves are possible, backtrack (return to the previous cell and try another direction).
// 	3.	Base Cases:
// 	•	If the current cell is out of bounds or a wall, return failure.
// 	•	If the current cell is the destination, return success.
// Directions for moving: down, up, right, left
var directions = [][]int{
	{1, 0},  // Down
	{-1, 0}, // Up
	{0, 1},  // Right
	{0, -1}, // Left
}

// SolveMaze finds a path in the maze from start to end using backtracking
func SolveMaze(maze [][]int, start, end []int) bool {
	rows := len(maze)
	cols := len(maze[0])

	// Helper function for recursion
	var backtrack func(x, y int) bool
	backtrack = func(x, y int) bool {
		// Base cases
		if x < 0 || y < 0 || x >= rows || y >= cols || maze[x][y] != 0 {
			return false
		}
		if x == end[0] && y == end[1] {
			return true // Destination reached
		}

		// Mark the cell as visited
		maze[x][y] = 2

		// Explore all four directions
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if backtrack(newX, newY) {
				return true
			}
		}

		// Backtrack: Unmark the cell
		maze[x][y] = 0
		return false
	}

	return backtrack(start[0], start[1])
}

func mainSolveMaze() {
	// Example maze: 0 = path, 1 = wall
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	start := []int{0, 0} // Top-left corner
	end := []int{4, 4}   // Bottom-right corner

	if SolveMaze(maze, start, end) {
		fmt.Println("Path found!")
	} else {
		fmt.Println("No path found.")
	}
}