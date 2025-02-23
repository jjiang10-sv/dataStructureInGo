package graph

import (
	"container/heap"
	"fmt"
	"math"
)

// heapData is an internal struct that implements the standard heap interface
// and keeps the data stored in the heap.
type heapData struct {
	queue []*Vertex
}

var (
	_ = heap.Interface(&heapData{}) // heapData is a standard heap
)

// Less compares two objects and returns true if the first one should go
// in front of the second one in the heap.
func (h *heapData) Less(i, j int) bool {
	if i > len(h.queue) || j > len(h.queue) {
		return false
	}
	return h.queue[i].Distance < h.queue[j].Distance
}

// Len returns the number of items in the Heap.
func (h *heapData) Len() int { return len(h.queue) }

// Swap implements swapping of two elements in the heap. This is a part of standard
// heap interface and should never be called directly.
func (h *heapData) Swap(i, j int) {
	h.queue[i], h.queue[j] = h.queue[j], h.queue[i]

}

// Push is supposed to be called by heap.Push only.
func (h *heapData) Push(kv interface{}) {
	keyValue := kv.(*Vertex)
	h.queue = append(h.queue, keyValue)
}

// Pop is supposed to be called by heap.Pop only.
func (h *heapData) Pop() interface{} {
	item := h.queue[len(h.queue)-1]
	h.queue = h.queue[0 : len(h.queue)-1]
	return item
}

func getShortestPath1(startNode *Node, endNode *Node, g *ItemGraph) ([]string, int) {
	visited := make(map[string]bool)
	//visited := make([]bool, len(g.Nodes))
	dist := make(map[string]int)
	prev := make(map[string]string)
	//pq := make(PriorityQueue, 1)
	//heap.Init(&pq)
	pq := heapData{}
	//pq := q.NewQ()
	start := Vertex{
		Node:     startNode,
		Distance: 0,
	}
	for _, nval := range g.Nodes {
		dist[nval.Value] = math.MaxInt64
	}
	dist[startNode.Value] = start.Distance
	heap.Push(&pq, &start)
	//im := 0
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Vertex)
		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true
		near := g.Edges[*v.Node]

		for _, edge := range near {
			if !visited[edge.Node.Value] {
				if dist[v.Node.Value]+edge.Weight < dist[edge.Node.Value] {
					updatedDist := dist[v.Node.Value] + edge.Weight
					store := Vertex{
						Node:     edge.Node,
						Distance: updatedDist,
					}
					dist[edge.Node.Value] = updatedDist
					//prev[edge.Node.Value] = fmt.Sprintf("->%s", v.Node.Value)
					prev[edge.Node.Value] = v.Node.Value
					heap.Push(&pq, &store)
				}
				
			}
		}
	}
	fmt.Println(dist)
	fmt.Println(prev)
	pathval := prev[endNode.Value]
	var finalArr []string
	finalArr = append(finalArr, endNode.Value)
	for pathval != startNode.Value {
		finalArr = append(finalArr, pathval)
		pathval = prev[pathval]
	}
	finalArr = append(finalArr, pathval)
	fmt.Println(finalArr)
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	return finalArr, dist[endNode.Value]

}
