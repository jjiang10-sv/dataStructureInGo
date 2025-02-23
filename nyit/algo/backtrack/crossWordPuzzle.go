package main

import (
	"fmt"
)

// Crossword puzzles are word games where players fill in grids with words or phrases based on clues. These puzzles are structured with a mix of blank and shaded squares, 
//where the goal is to write correct answers in the blank spaces. The answers must fit according 
//to their length and intersect correctly with other words in the puzzle.

// Key Features of Crossword Puzzles:

// 	1.	Grid Structure: Typically rectangular, divided into squares for letters. Blank squares are filled with letters, and shaded squares separate words or phrases.
// 	2.	Clues: Provided as definitions, synonyms, or riddles. Clues may be straightforward, cryptic, or themed.
// 	•	Across: Words written horizontally.
// 	•	Down: Words written vertically.
// 	3.	Intersections: Letters in one word often overlap with letters from intersecting words, providing hints and constraints.
// 	4.	Difficulty: Varies based on clue style, word length, and complexity of the theme.

// Types of Crossword Puzzles:

// 	1.	Standard Crosswords: Use direct clues with dictionary-style definitions.
// 	2.	Cryptic Crosswords: Clues involve wordplay, anagrams, or double meanings.
// 	3.	Themed Crosswords: A central theme connects many answers.
// 	4.	American vs. British Styles:
// 	•	American Crosswords: Symmetrical grids, often with few black squares.
// 	•	British Crosswords: May feature more irregular grid designs and cryptic clues.

// Benefits of Solving Crosswords:

// 	1.	Cognitive Stimulation: Enhances vocabulary, general knowledge, and pattern recognition.
// 	2.	Stress Relief: Offers a focused, meditative activity.
// 	3.	Entertainment: A satisfying challenge for puzzle enthusiasts.

// Tools to Create and Solve Crosswords:

// 	•	Online Generators: Tools like Crossword Compiler or EclipseCrossword.
// 	•	Solving Apps: Apps such as “NYT Crossword” or “Crossword Solver.”

// If you’re looking for custom puzzles or resources to practice, let me know, and I can assist further!
const EMPTY = '-'

// Grid represents the crossword grid
type Grid [][]rune

// isSafeCross checks if a word can be placed at a given position
func isSafeCross(grid Grid, word string, row, col int, isHorizontal bool) bool {
	for i, ch := range word {
		if isHorizontal {
			if col+i >= len(grid[row]) || (grid[row][col+i] != EMPTY && grid[row][col+i] != ch) {
				return false
			}
		} else {
			if row+i >= len(grid) || (grid[row+i][col] != EMPTY && grid[row+i][col] != ch) {
				return false
			}
		}
	}
	return true
}

// placeWord places a word in the grid
func placeWord(grid Grid, word string, row, col int, isHorizontal bool) []int {
	original := []int{}
	for i, ch := range word {
		if isHorizontal {
			if grid[row][col+i] == EMPTY {
				original = append(original, col+i)
				grid[row][col+i] = ch
			}
		} else {
			if grid[row+i][col] == EMPTY {
				original = append(original, row+i)
				grid[row+i][col] = ch
			}
		}
	}
	return original
}

// removeWord restores the grid to its original state
func removeWord(grid Grid, original []int, row, col int, isHorizontal bool) {
	for _, index := range original {
		if isHorizontal {
			grid[row][index] = EMPTY
		} else {
			grid[index][col] = EMPTY
		}
	}
}

// solveCrossword solves the crossword puzzle using backtracking
func solveCrossword(grid Grid, words []string, index int) bool {
	if index == len(words) {
		return true // All words are placed
	}

	word := words[index]
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if isSafeCross(grid, word, row, col, true) { // Try horizontal
				original := placeWord(grid, word, row, col, true)
				if solveCrossword(grid, words, index+1) {
					return true
				}
				removeWord(grid, original, row, col, true)
			}
			if isSafeCross(grid, word, row, col, false) { // Try vertical
				original := placeWord(grid, word, row, col, false)
				if solveCrossword(grid, words, index+1) {
					return true
				}
				removeWord(grid, original, row, col, false)
			}
		}
	}
	return false
}

// printGrid prints the crossword grid
func printGrid(grid Grid) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func mainCross() {
	// Example crossword grid and words
	grid := Grid{
		[]rune("----"),
		[]rune("----"),
		[]rune("----"),
		[]rune("----"),
	}
	words := []string{"this", "is", "fun"}

	if solveCrossword(grid, words, 0) {
		fmt.Println("Solution:")
		printGrid(grid)
	} else {
		fmt.Println("No solution exists!")
	}
}
