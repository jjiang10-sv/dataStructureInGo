package main

import "fmt"

// •	The function LongestCommonSubsequence takes two strings as input and computes the length of the LCS.
// •	A 2D slice (dp) is used to store the results of subproblems.
// •	The algorithm iteratively builds the solution bottom-up, filling the dp table based on matches between characters of the input strings.

// How It Works:

// 	1.	Initialization: A 2D array dp of size (n+1) x (m+1) is initialized to store the length of LCS for substrings.
// 	2.	Filling the DP Table:
// 	•	If s1[i-1] == s2[j-1], the LCS length increases by 1 (dp[i][j] = dp[i-1][j-1] + 1).
// 	•	Otherwise, it takes the maximum value from the adjacent subproblems (dp[i][j] = max(dp[i-1][j], dp[i][j-1])).
// 	3.	Reconstruction: The LCS string is reconstructed by tracing the path from dp[n][m] backward.

// LongestCommonSubsequence computes the LCS of two strings
func LongestCommonSubsequence(s1, s2 string) string {
    n, m := len(s1), len(s2)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }

    // Build the dp table
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    // Reconstruct the LCS
    lcs := []byte{}
    i, j := n, m
    for i > 0 && j > 0 {
        if s1[i-1] == s2[j-1] {
            lcs = append([]byte{s1[i-1]}, lcs...)
            i--
            j--
        } else if dp[i-1][j] > dp[i][j-1] {
            i--
        } else {
            j--
        }
    }

    return string(lcs)
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    s1 := "ABCBDAB"
    s2 := "BDCAB"
    fmt.Println("Longest Common Subsequence:", LongestCommonSubsequence(s1, s2))
}