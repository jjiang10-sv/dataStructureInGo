package main

import (
	"fmt"
)

func solveCryptarithmeticba() {
	letters := []rune{'S', 'E', 'N', 'D', 'M', 'O', 'R', 'Y'}
	digits := make([]bool, 10) // To track which digits are used
	mapping := make(map[rune]int)

	if backtrack(letters, digits, mapping, 0) {
		fmt.Println("Solution found:")
		for letter, digit := range mapping {
			fmt.Printf("%c = %d\n", letter, digit)
		}
	} else {
		fmt.Println("No solution exists.")
	}
}

func backtrack(letters []rune, digits []bool, mapping map[rune]int, index int) bool {
	if index == len(letters) {
		// All letters are assigned; check the equation
		return isValid(mapping)
	}

	// Try assigning each unused digit to the current letter
	for d := 0; d <= 9; d++ {
		if !digits[d] { // Check if the digit is unused
			// Assign digit to the current letter
			mapping[letters[index]] = d
			digits[d] = true

			// Recur to assign digits to the next letter
			if backtrack(letters, digits, mapping, index+1) {
				return true
			}

			// Backtrack
			delete(mapping, letters[index])
			digits[d] = false
		}
	}
	return false
}

func isValid(mapping map[rune]int) bool {
	// Convert SEND, MORE, MONEY to numbers using the mapping
	send := getValue("SEND", mapping)
	more := getValue("MORE", mapping)
	money := getValue("MONEY", mapping)

	// Check if SEND + MORE = MONEY
	return send+more == money && mapping['S'] != 0 && mapping['M'] != 0
}

func getValue(word string, mapping map[rune]int) int {
	value := 0
	for _, letter := range word {
		value = value*10 + mapping[letter]
	}
	return value
}

func mainCba() {
	solveCryptarithmetic()
}

// Explanation

// 	1.	Input Letters: The problem involves eight unique letters: S, E, N, D, M, O, R, and Y.
// 	2.	Backtracking Function:
// 	•	Assign a digit to each letter.
// 	•	Check if the assignment satisfies the constraints (e.g., SEND + MORE = MONEY).
// 	•	If it doesn’t, undo the assignment (backtrack) and try another digit.
// 	3.	Validation:
// 	•	Convert the words SEND, MORE, and MONEY into their numeric equivalents using the current mapping.
// 	•	Ensure that the first letters (S and M) are non-zero.
// 	4.	Base Case:
// 	•	When all letters have been assigned, validate the equation.
