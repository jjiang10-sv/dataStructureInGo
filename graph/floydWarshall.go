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

func main1() {
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

