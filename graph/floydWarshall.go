package graph

import (
	"fmt"
	"math"
)

const inf = math.MaxInt32 // Representing infinity

func floydWarshall1(vertices int, graph [][]int) [][]int {
	// Initialize distance matrix with input graph values
	dist := make([][]int, vertices)
	for i := range dist {
		dist[i] = make([]int, vertices)
		for j := range dist[i] {
			dist[i][j] = graph[i][j]
		}
	}

	// Floyd-Warshall algorithm
	// recursive approach and dynamic programming; k as the mediate value
	for k := 0; k < vertices; k++ {
		// iterate the vertices because it is an all pair shortest path
		for i := 0; i < vertices; i++ {
			// skip the row k. no need to update on it.
			if i != k {
				// iterate and update the distance
				for j := 0; j < vertices; j++ {
					// no need to update if j is equal to the mediate value
					if j != k {
						// If i to k and k to j are reachable
						// if there is a edge between i and k also and k and j; then check if the distance is less than the current distance
						if dist[i][k] != inf && dist[k][j] != inf && dist[i][j] > dist[i][k]+dist[k][j] {
							dist[i][j] = dist[i][k] + dist[k][j]
						}
					}

				}
			}

		}
	}

	// Detect negative cycles
	for i := 0; i < vertices; i++ {
		if dist[i][i] < 0 {
			fmt.Println("Graph contains a negative weight cycle")
			return nil
		}
	}

	return dist
}

func allPairShortestPath() {
	vertices := 4
	graph := [][]int{
		{0, 3, inf, 5},
		{2, 0, inf, 4},
		{inf, 1, 0, inf},
		{inf, inf, 2, 0},
	}

	dist := floydWarshall1(vertices, graph)

	if dist != nil {
		fmt.Println("Shortest distances between every pair of vertices:")
		for i := 0; i < vertices; i++ {
			for j := 0; j < vertices; j++ {
				if dist[i][j] == inf {
					fmt.Print("inf ")
				} else {
					fmt.Printf("%3d ", dist[i][j])
				}
			}
			fmt.Println()
		}
	}
}
