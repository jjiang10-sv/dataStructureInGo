package main

import "fmt"

// The Map Coloring Problem involves assigning colors to regions on a map such that no two adjacent regions share the same color. This is a classic graph coloring problem, where the map is represented as a graph, and the regions are nodes connected by edges.

// Here’s how you can solve it using backtracking:

// Algorithm Overview

// 	1.	Graph Representation:
// 	•	Represent the map as a graph  G(V, E) , where:
// 	•	 V  is the set of nodes (regions on the map).
// 	•	 E  is the set of edges (adjacent regions).
// 	•	Use an adjacency matrix or list for implementation.
// 	2.	Backtracking Approach:
// 	•	Start with the first node and try assigning a color.
// 	•	Move to the next node and try to assign a valid color (different from its adjacent nodes).
// 	•	If a valid assignment is not possible, backtrack to the previous node and try a different color.
// 	3.	Base Cases:
// 	•	If all nodes are colored, return success.
// 	•	If no valid color exists for a node, backtrack.

// Go Implementation

// Below is an implementation of the map coloring problem in Go:

func isSafe(node int, graph [][]int, colors []int, color int) bool {
	for adj := 0; adj < len(graph); adj++ {
		if graph[node][adj] == 1 && colors[adj] == color {
			return false
		}
	}
	return true
}

func colorMap(node int, graph [][]int, m int, colors []int) bool {
	// Base case: If all nodes are colored
	if node == len(graph) {
		return true
	}

	// Try assigning colors
	for c := 1; c <= m; c++ {
		if isSafe(node, graph, colors, c) {
			colors[node] = c // Assign the color

			// Recurse to color the next node
			if colorMap(node+1, graph, m, colors) {
				return true
			}

			// Backtrack: Remove the color assignment
			colors[node] = 0
		}
	}
	return false
}

func mainMapColor() {
	// Example adjacency matrix for the graph
	graph := [][]int{
		{0, 1, 1, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 1},
		{1, 0, 1, 0},
	}
	numNodes := len(graph)

	m := 3                     // Number of colors
	colors := make([]int, numNodes) // Array to store colors for each node

	if colorMap(0, graph, m, colors) {
		fmt.Println("Solution found:", colors)
	} else {
		fmt.Println("No solution exists")
	}
}