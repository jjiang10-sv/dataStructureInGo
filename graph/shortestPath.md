Here’s how you can implement the Bellman-Ford algorithm in Go to find the shortest paths from a single source vertex to all other vertices in a weighted graph. Bellman-Ford is particularly useful for graphs that contain negative weight edges and can detect negative weight cycles.

Bellman-Ford Algorithm in Go

	1.	Define the Graph Structure: Represent the graph using edges and weights.
	2.	Initialize Distances: Start distances from the source node with 0 and set all others to infinity.
	3.	Relax Edges: Iterate V-1 times (where V is the number of vertices) to relax edges and update distances.
	4.	Detect Negative Cycles: Do one more iteration to check for any distance reduction; if detected, a negative cycle exists.

Code Example

package main

import (
    "fmt"
    "math"
)

type Edge struct {
    source, destination, weight int
}

// BellmanFord finds shortest paths from src to all vertices.
// Returns an error if a negative weight cycle is detected.
func BellmanFord(vertices int, edges []Edge, src int) ([]int, error) {
    // Initialize distances from src to all other vertices as infinity
    distances := make([]int, vertices)
    for i := range distances {
        distances[i] = math.MaxInt32
    }
    distances[src] = 0

    // Relax edges |V| - 1 times
    for i := 0; i < vertices-1; i++ {
        for _, edge := range edges {
            if distances[edge.source] != math.MaxInt32 && distances[edge.source]+edge.weight < distances[edge.destination] {
                distances[edge.destination] = distances[edge.source] + edge.weight
            }
        }
    }

    // Check for negative-weight cycles
    for _, edge := range edges {
        if distances[edge.source] != math.MaxInt32 && distances[edge.source]+edge.weight < distances[edge.destination] {
            return nil, fmt.Errorf("graph contains a negative weight cycle")
        }
    }

    return distances, nil
}

func main() {
    vertices := 5
    edges := []Edge{
        {0, 1, -1},
        {0, 2, 4},
        {1, 2, 3},
        {1, 3, 2},
        {1, 4, 2},
        {3, 2, 5},
        {3, 1, 1},
        {4, 3, -3},
    }

    // Run Bellman-Ford from vertex 0
    distances, err := BellmanFord(vertices, edges, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Vertex distances from source:")
        for i, d := range distances {
            fmt.Printf("Distance to vertex %d: %d\n", i, d)
        }
    }
}

Explanation of the Code

	•	Edge Struct: This struct represents an edge with a source, destination, and weight.
	•	BellmanFord Function: The function takes the number of vertices, an array of edges, and the source vertex.
	•	It initializes the distance to all vertices as infinity (math.MaxInt32) except for the source.
	•	Then, it relaxes each edge V-1 times.
	•	Finally, it checks for negative weight cycles.
	•	Main Function: Sets up a graph and calls the BellmanFord function with the source vertex.

Example Output

For the provided example with src = 0, the output would be:

Vertex distances from source:
Distance to vertex 0: 0
Distance to vertex 1: -1
Distance to vertex 2: 2
Distance to vertex 3: -2
Distance to vertex 4: 1

Complexity

	•	Time Complexity: ￼, where V is the number of vertices and E is the number of edges.
	•	Space Complexity: ￼ due to the storage of distances.


The Floyd-Warshall algorithm is an all-pairs shortest path algorithm, which finds the shortest paths between all pairs of nodes in a weighted graph. It’s particularly useful for dense graphs or when you need the shortest paths between every pair of vertices.

Floyd-Warshall Algorithm Steps

	1.	Initialize the Distance Matrix: Use a matrix where dist[i][j] represents the distance from vertex i to vertex j.
	2.	Set Distances:
	•	For each edge (i, j) with weight w, set dist[i][j] = w.
	•	Set dist[i][i] = 0 for each vertex i.
	•	For vertices not directly connected, initialize their distance to infinity.
	3.	Relax All Pairs: For each vertex k, update all pairs (i, j) by setting dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j]).
	4.	Negative Cycles: If any dist[i][i] becomes negative, the graph contains a negative weight cycle.

Code Example

Here’s how you can implement the Floyd-Warshall algorithm in Go:

package main

import (
    "fmt"
    "math"
)

const inf = math.MaxInt32 // Representing infinity

func floydWarshall(vertices int, graph [][]int) [][]int {
    // Initialize distance matrix with input graph values
    dist := make([][]int, vertices)
    for i := range dist {
        dist[i] = make([]int, vertices)
        for j := range dist[i] {
            dist[i][j] = graph[i][j]
        }
    }

    // Floyd-Warshall algorithm
    for k := 0; k < vertices; k++ {
        for i := 0; i < vertices; i++ {
            for j := 0; j < vertices; j++ {
                // If i to k and k to j are reachable
                if dist[i][k] != inf && dist[k][j] != inf && dist[i][j] > dist[i][k]+dist[k][j] {
                    dist[i][j] = dist[i][k] + dist[k][j]
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

func main() {
    vertices := 4
    graph := [][]int{
        {0, 3, inf, 5},
        {2, 0, inf, 4},
        {inf, 1, 0, inf},
        {inf, inf, 2, 0},
    }

    dist := floydWarshall(vertices, graph)

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

Explanation of the Code

	•	Initialize the Distance Matrix: We create a matrix dist from the input graph, where graph[i][j] is the weight of the edge from vertex i to vertex j.
	•	Floyd-Warshall Algorithm Loop:
	•	For each vertex k, it iterates over every pair of vertices (i, j) and updates the distance if dist[i][j] > dist[i][k] + dist[k][j].
	•	Negative Cycle Detection: If any diagonal element dist[i][i] is negative, the graph contains a negative cycle.
	•	Output: The distance matrix dist is printed, showing the shortest path between every pair of vertices.

Example Output

For the input graph, the output would be:

Shortest distances between every pair of vertices:
  0   3   7   5 
  2   0   6   4 
  3   1   0   5 
  5   3   2   0 

Complexity

	•	Time Complexity: ￼, where V is the number of vertices.
	•	Space Complexity: ￼, since it stores all-pairs distances in a matrix.

In the Bellman-Ford algorithm, we require |V| - 1 iterations over all edges to guarantee that we find the shortest paths in a graph with |V| vertices. Each iteration ensures that distance values propagate through the graph, accounting for the paths created by edges.

Why |V| - 1 Iterations?

This is because a shortest path in a graph with no cycles has at most |V| - 1 edges, where |V| is the number of vertices. For a graph with edges forming paths of up to |V| - 1 steps, the Bellman-Ford algorithm iteratively updates the shortest path estimate for each vertex, ensuring that each possible vertex-to-vertex path is accounted for.

Handling Negative Cycles

If a graph contains a negative cycle, additional iterations are required to propagate the “minus infinity” values from this cycle through the graph. After the initial |V| - 1 iterations, any further updates indicate the presence of a negative-weight cycle, as they continue decreasing the shortest path estimates beyond the feasible |V| - 1 steps. In this case:

	•	The algorithm detects these additional changes to recognize the negative cycles.
	•	This helps identify which vertices and edges are affected by the cycle, as any reachable vertices continue to have their distances lowered.