// The Rod Cutting Problem is a classic optimization problem that can be solved using dynamic programming. The problem is defined as follows:

// Problem Definition:

// You are given a rod of length n and a set of prices for each possible length of the rod (from 1 to n). The task is to determine the maximum revenue you can obtain by cutting the rod into smaller pieces (if necessary) and selling them. You can cut the rod into pieces of any length, and the goal is to maximize the revenue.

// Approach:

// 	1.	Recursive Solution: A naive approach would be to recursively try every possible way to cut the rod, calculating the revenue for each possible cut. However, this approach results in redundant calculations, which can be optimized.
// 	2.	Dynamic Programming Solution: The problem has an optimal substructure, meaning that the optimal solution to the problem can be derived from optimal solutions to smaller subproblems. Thus, dynamic programming is an ideal approach.

// Dynamic Programming Approach:

// We can define a table dp[] where dp[i] represents the maximum revenue obtainable by cutting a rod of length i. The recurrence relation is:

// dp[i] = max(price[j] + dp[i-j-1] for all 0 <= j < i)

// Here, price[j] is the price of a rod of length j+1, and dp[i-j-1] is the maximum revenue obtainable from the remaining length of the rod.

// Solution in Go:

package main

import "fmt"

// RodCutting solves the rod cutting problem using dynamic programming.
func RodCutting(prices []int, n int) int {
	// dp[i] will store the maximum value obtainable for a rod of length i
	dp := make([]int, n+1)

	// Build the dp array in a bottom-up manner
	for i := 1; i <= n; i++ {
		// Try cutting the rod at every possible length j
		for j := 0; j < i; j++ {
			dp[i] = max(dp[i], prices[j] + dp[i-j-1])
		}
	}

	// The maximum revenue for the rod of length n is stored in dp[n]
	return dp[n]
}


func mainRod() {
	// Prices for different lengths of the rod
	// For example, prices[i] is the price of a rod of length i+1
	prices := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	n := 8

	// Call the RodCutting function
	maxRevenue := RodCutting(prices, n)
	fmt.Printf("Maximum revenue obtainable: %d\n", maxRevenue)
}

// Explanation:

// 	1.	Input:
// 	•	prices is an array where prices[i] represents the price of a rod of length i+1.
// 	•	n is the length of the rod.
// 	2.	Dynamic Programming Table:
// 	•	dp[i] represents the maximum revenue obtainable from a rod of length i.
// 	3.	Recurrence Relation:
// 	•	For each rod length i, we try every possible cut j and calculate the maximum revenue by considering the price of cutting a rod of length j+1 and the remaining rod length i-j-1.
// 	4.	Output:
// 	•	The maximum revenue obtainable for a rod of length n is found in dp[n].

// Example Output:

// For the input:

// prices := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
// n := 8

// The output would be:

// Maximum revenue obtainable: 22

// Time Complexity:

// 	•	Time Complexity: ￼, where ￼ is the length of the rod. The outer loop iterates over all lengths from 1 to n, and the inner loop computes the maximum value by checking all possible cuts.
// 	•	Space Complexity: ￼ for the dp array.

// Conclusion:

// The Rod Cutting problem is a classic example of dynamic programming that optimally solves the problem by breaking it down into subproblems. This approach avoids redundant calculations, making it efficient even for larger inputs.

// For more information, see references like Introduction to Algorithms by Cormen et al. or other resources on dynamic programming.