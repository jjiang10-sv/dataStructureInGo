package main

import (
	"fmt"
	"math"
)

// Function to find the minimum cost of multiplying matrices
func matrixChainOrder(p []int, n int) int {
	// Create a 2D array to store the minimum multiplication costs
	// m[i][j] will store the minimum number of scalar multiplications needed to multiply matrices Ai..Aj
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}

	// cost is zero when multiplying one matrix
	for i := 1; i < n; i++ {
		m[i][i] = 0
	}

	// l is the chain length
	for l := 2; l < n; l++ {
		for i := 1; i < n-l+1; i++ {
			j := i + l - 1
			m[i][j] = math.MaxInt
			// Try every possible split point k between i and j
			for k := i; k < j; k++ {
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				if q < m[i][j] {
					m[i][j] = q
				}
			}
		}
	}

	// Return the minimum cost to multiply the entire matrix chain
	return m[1][n-1]
}

func mainChina() {
	// Dimensions of the matrices
	// Matrices A1, A2, A3, ..., An
	// Matrix Ai has dimensions p[i-1] x p[i]
	p := []int{10, 20, 30, 40, 30}
	n := len(p)

	// Call the matrixChainOrder function
	minCost := matrixChainOrder(p, n)
	fmt.Printf("Minimum number of multiplications is %d\n", minCost)
}

// The Matrix Chain Multiplication Problem is a classic problem in dynamic programming. It involves finding the most efficient way to multiply a chain of matrices. Given a sequence of matrices, the goal is to determine the optimal parenthesization of these matrices in order to minimize the number of scalar multiplications required.

// Problem Definition:

// Given a sequence of matrices A1, A2, ..., An, the dimensions of matrix Ai are given by p[i-1] x p[i]. The goal is to determine the most efficient way to multiply these matrices together, which can be solved by dynamic programming.

// The problem is solved by determining the minimum number of scalar multiplications needed to compute the product of matrices in the chain.

// Dynamic Programming Approach:

// 	1.	Matrix Chain Order: We define a table m[i][j] which stores the minimum number of scalar multiplications required to multiply matrices Ai through Aj.
// 	2.	Recursion: For any matrix chain from i to j, we compute the cost of splitting the chain at every possible position k (where i <= k < j), and then use the result from smaller subproblems.
// 	3.	Optimal Substructure: The problem has an optimal substructure because the cost to multiply matrices from i to j depends on the costs of multiplying matrices from i to k and from k+1 to j.

// // Explanation of Code:

// // 	1.	Input: The p array represents the dimensions of the matrices in the chain. For example, if p = [10, 20, 30, 40, 30], it means there are 4 matrices:
// // 	•	A1 is of size 10 x 20
// // 	•	A2 is of size 20 x 30
// // 	•	A3 is of size 30 x 40
// // 	•	A4 is of size 40 x 30
// // 	2.	Dynamic Programming Table:
// // 	•	The m table is a 2D array where m[i][j] represents the minimum cost to multiply matrices from Ai to Aj.
// // 	•	We initialize the diagonal of the matrix m[i][i] = 0 because no multiplication is needed when there’s only one matrix.
// // 	3.	Filling the DP Table:
// // 	•	We loop over increasing lengths of the matrix chain (l), starting from 2 (two matrices) up to n (the entire chain).
// // 	•	For each chain of length l, we try all possible places to split the chain (k), and we calculate the cost of multiplying the two resulting subchains.
// // 	•	The recursive formula for the minimum cost is:
// // m[i][j] = min(m[i][k] + m[k+1][j] + p[i-1] * p[k] * p[j]) where i <= k < j.
// // 	4.	Result: The minimum number of scalar multiplications required for the entire matrix chain multiplication is stored in m[1][n-1].

// // Example Output:

// // For the input p = [10, 20, 30, 40, 30], the output will be:

// // Minimum number of multiplications is 30000

// // Time Complexity:

// // 	•	Time Complexity: ￼, where ￼ is the number of matrices. This is due to the three nested loops: one for the chain length, one for the starting point of the subchain, and one for the split point.
// // 	•	Space Complexity: ￼ for storing the DP table.

// // Conclusion:

// // The Matrix Chain Multiplication problem is a classic example of dynamic programming where the goal is to minimize the computational cost by breaking the problem down into smaller subproblems. The algorithm described efficiently solves this problem in ￼ time, making it suitable for moderate-sized inputs.

// // For further reading, you can check out resources like Introduction to Algorithms by Cormen et al., which covers dynamic programming techniques in great detail.
