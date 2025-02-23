package main

import (
	"fmt"
)

// We use a dynamic programming (DP) table where dp[i][j] represents the minimum number of operations needed to convert the first i characters of string word1 to the first j characters of string word2.

// Algorithm Steps

// 	1.	Base Cases:
// 	•	If one string is empty, the edit distance equals the length of the other string (i.e., all characters must be inserted or removed).
// 	2.	Recursive Relation:
// 	•	If the characters match (word1[i-1] == word2[j-1]), no new operation is needed:

// dp[i][j] = dp[i-1][j-1]

// 	•	If the characters don’t match, consider the minimum of three possible operations:
// 	•	Insert a character: dp[i][j-1] + 1
// 	•	Remove a character: dp[i-1][j] + 1
// 	•	Replace a character: dp[i-1][j-1] + 1
// 	3.	Iterative Calculation:
// 	•	Fill the DP table bottom-up, using the above relations.
// 	4.	Result:
// 	•	The final edit distance is stored in dp[len(word1)][len(word2)].

// Min function to find the smallest among three numbers
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// EditDistance calculates the minimum edit distance between two strings
func EditDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)

	// Create a 2D DP array
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i // Deleting all characters from word1
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // Inserting all characters to word1
	}

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // Characters match
			} else {
				dp[i][j] = min(
					dp[i-1][j],   // Remove
					dp[i][j-1],   // Insert
					dp[i-1][j-1], // Replace
				) + 1
			}
		}
	}

	// The answer is in dp[m][n]
	return dp[m][n]
}

func allPairShortestPath() {
	word1 := "kitten"
	word2 := "sitting"
	fmt.Printf("Edit Distance between '%s' and '%s': %d\n", word1, word2, EditDistance(word1, word2))
}
