package main

import (
	"fmt"
)

// Graph structure
type Graph struct {
	vertices int
	adjList  map[int][]int
	inDegree []int
}

// NewGraph initializes a graph
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
		inDegree: make([]int, vertices),
	}
}

// AddEdge adds a directed edge (u -> v)
func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.inDegree[v]++ // Increase in-degree of destination node
}

// Kahn's Algorithm (BFS-based Topological Sort)
func (g *Graph) TopologicalSort() []int {
	queue := []int{}
	order := []int{}

	// Push nodes with in-degree = 0 to queue
	for i := 0; i < g.vertices; i++ {
		if g.inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Process the queue
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // Dequeue
		order = append(order, node)

		// Reduce in-degree of neighbors
		for _, neighbor := range g.adjList[node] {
			g.inDegree[neighbor]--
			if g.inDegree[neighbor] == 0 {
				queue = append(queue, neighbor) // Enqueue if in-degree becomes 0
			}
		}
	}

	// If topological order contains all nodes, return it
	if len(order) == g.vertices {
		return order
	}

	return []int{} // Return empty if cycle detected (not a DAG)
}

// Main function
func main() {
	g := NewGraph(6)
	g.AddEdge(5, 2)
	g.AddEdge(5, 0)
	g.AddEdge(4, 0)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)

	fmt.Println("Topological Sort:", g.TopologicalSort())
}
