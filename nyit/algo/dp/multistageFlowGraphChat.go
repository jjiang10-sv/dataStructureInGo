// Here is a more complex input graph for the multi-stage graph problem. This graph includes 8 nodes, and it is divided into 4 stages, with varying edge weights.

// Updated Example:

// Graph Representation

// Stage 1:    (1)
//            / |  \
// Stage 2: (2) (3) (4)
//            \   |   /
// Stage 3:    (5) (6)
//               |  /
// Stage 4:     (7)

// Adjacency Matrix Representation

// [
//   [0, 2, 3, 6, 0, 0, 0, 0],  // Node 1
//   [0, 0, 0, 0, 5, 8, 0, 0],  // Node 2
//   [0, 0, 0, 0, 7, 3, 0, 0],  // Node 3
//   [0, 0, 0, 0, 4, 6, 0, 0],  // Node 4
//   [0, 0, 0, 0, 0, 0, 2, 4],  // Node 5
//   [0, 0, 0, 0, 0, 0, 6, 3],  // Node 6
//   [0, 0, 0, 0, 0, 0, 0, 2],  // Node 7
//   [0, 0, 0, 0, 0, 0, 0, 0],  // Node 8
// ]

// Updated Code Implementation

package main

import (
	"fmt"
	"math"
)


// Find the shortest path in a multi-stage graph
func shortestPathMultiStage(graph Graph, stages int) int {
	// Get the number of nodes
	n := len(graph.edges)

	// Create a distance array and initialize it to infinity
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	// Initialize the distance to the sink (last node) as 0
	dist[n-1] = 0

	// Traverse the graph from the second last node back to the source
	for i := n - 2; i >= 0; i-- {
		// Check all possible connections from node `i`
		for j := i + 1; j < n; j++ {
			if graph.edges[i][j] > 0 { // Edge exists
				dist[i] = minMC(dist[i], graph.edges[i][j]+dist[j])
			}
		}
	}

	return dist[0] // Return the distance to the source
}

// Utility function to find the minimum of two integers
func minMC(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mainMC() {
	// Example multi-stage graph with 8 nodes and 4 stages
	graph := Graph{
		edges: [][]int{
			{0, 2, 3, 6, 0, 0, 0, 0},
			{0, 0, 0, 0, 5, 8, 0, 0},
			{0, 0, 0, 0, 7, 3, 0, 0},
			{0, 0, 0, 0, 4, 6, 0, 0},
			{0, 0, 0, 0, 0, 0, 2, 4},
			{0, 0, 0, 0, 0, 0, 6, 3},
			{0, 0, 0, 0, 0, 0, 0, 2},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	// Find the shortest path
	stages := 4
	result := shortestPathMultiStage(graph, stages)

	fmt.Printf("The shortest path distance is: %d\n", result)
}

// Output:

// For the updated graph, the program calculates the shortest path from node 1 (source) to node 8 (sink). The output is:

// The shortest path distance is: 9

// This graph and code demonstrate how to handle more complex multi-stage graphs with a larger number of nodes and varying paths between them.
