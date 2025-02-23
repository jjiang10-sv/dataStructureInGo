package main

import (
	"fmt"
)

// Explanation of the Code

// 	1.	Word to Number Conversion: The wordToNumber function converts a string of letters into a number based on the current letter-to-digit mapping.
// 	2.	Permutation of Digits: The permute function generates all possible digit assignments (0-9) for the unique letters in the problem. For each assignment, it checks if the equation holds true.
// 	3.	Backtracking: The permute function is implemented using backtracking. It tries all possible combinations of digits for each letter and evaluates the equation SEND + MORE == MONEY.
// 	4.	Brute Force Approach: This is a brute-force approach that checks all possible digit assignments for the letters. It is not efficient for large problems but works for small ones like this example.
// 	5.	Removing Duplicates: The removeDuplicates function ensures that only the unique letters are considered, avoiding unnecessary computations.
// 	6.	Solution Output: If a valid solution is found, it prints the letter-to-digit mapping.

// Helper function to convert a word into a number based on the letter-to-digit mapping
func wordToNumber(word string, letterToDigit map[rune]int) int {
	number := 0
	for _, letter := range word {
		number = number*10 + letterToDigit[letter]
	}
	return number
}

// Function to solve the Cryptarithmetic problem
func solveCryptarithmetic() {
	letters := "SENDMOREMONEY"
	uniqueLetters := removeDuplicates(letters)
	// Generate all possible digit combinations (0-9) for the unique letters
	var digits [10]int
	// Try all permutations of digits for the unique letters
	for i := 0; i < 10; i++ {
		digits[i] = i
	}

	// Brute force all combinations of digits for the letters
	permute(uniqueLetters, digits, func(letterToDigit map[rune]int) bool {
		send := wordToNumber("SEND", letterToDigit)
		more := wordToNumber("MORE", letterToDigit)
		money := wordToNumber("MONEY", letterToDigit)

		// Check if the equation SEND + MORE == MONEY is satisfied
		if send+more == money {
			// Print the solution
			fmt.Println("Solution Found:")
			for _, letter := range uniqueLetters {
				fmt.Printf("%c: %d\n", letter, letterToDigit[letter])
			}
			return true
		}
		return false
	})
}

// Helper function to remove duplicate letters from the string
func removeDuplicates(s string) []rune {
	letterSet := make(map[rune]struct{})
	var uniqueLetters []rune
	for _, letter := range s {
		if _, exists := letterSet[letter]; !exists {
			letterSet[letter] = struct{}{}
			uniqueLetters = append(uniqueLetters, letter)
		}
	}
	return uniqueLetters
}

// Permutation generator function
func permute(letters []rune, digits [10]int, callback func(map[rune]int) bool) {
	// Generate all permutations of the digits for the letters
	var permuteHelper func([]int, int)
	permuteHelper = func(digits []int, start int) {
		if start == len(digits) {
			// Map the letters to the current permutation of digits
			letterToDigit := make(map[rune]int)
			for i, letter := range letters {
				letterToDigit[letter] = digits[i]
			}
			if callback(letterToDigit) {
				return
			}
			return
		}
		for i := start; i < len(digits); i++ {
			digits[start], digits[i] = digits[i], digits[start]
			permuteHelper(digits, start+1)
			digits[start], digits[i] = digits[i], digits[start]
		}
	}
	permuteHelper(digits[:], 0)
}

func mainCry() {
	solveCryptarithmetic()
}
