package main

import "fmt"

// The 0/1 Knapsack Problem is a classic optimization problem where you have a knapsack with a fixed capacity and a set of items, each with a weight and a value. The goal is to maximize the total value of items placed in the knapsack without exceeding its capacity. The “0/1” indicates that each item can either be taken (1) or not taken (0).

// Dynamic Programming Approach

// The dynamic programming (DP) solution involves filling a table where each entry represents the maximum value that can be achieved with a given weight capacity and subset of items.

// Algorithm Steps

// 	1.	Define DP Table:
// Let dp[i][w] represent the maximum value that can be obtained using the first i items with a capacity of w.
// 	2.	Base Cases:
// 	•	dp[0][w] = 0 for all w (no items yield zero value).
// 	•	dp[i][0] = 0 for all i (zero capacity yields zero value).
// 	3.	Recursive Relation:
// 	•	If the weight of the current item i exceeds the current capacity w, the item cannot be included:

// dp[i][w] = dp[i-1][w]

// 	•	Otherwise, take the maximum of two choices:
// 	1.	Not taking the item: dp[i-1][w]
// 	2.	Taking the item: value[i-1] + dp[i-1][w - weight[i-1]]

// dp[i][w] = \max(dp[i-1][w], value[i-1] + dp[i-1][w - weight[i-1]])

// 	4.	Result:
// The result is stored in dp[n][W], where n is the number of items and W is the capacity of the knapsack.

// Knapsack solves the 0/1 Knapsack problem using dynamic programming
func Knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)
	// Create a 2D DP array
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Fill the DP table
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weights[i-1] > w {
				// Current item can't be included
				dp[i][w] = dp[i-1][w]
			} else {
				// Take the maximum of including or excluding the item
				dp[i][w] = max(dp[i-1][w], values[i-1]+dp[i-1][w-weights[i-1]])
			}
		}
	}

	// Return the maximum value
	return dp[n][capacity]
}

// // Helper function to find the maximum of two integers
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func mainq() {
	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	capacity := 7

	fmt.Printf("Maximum value in knapsack: %d\n", Knapsack(weights, values, capacity))
}