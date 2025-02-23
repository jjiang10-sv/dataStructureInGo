package graph

import (
	"fmt"
	"math"
	"sync"
)

// Enqueue adds an Node to the end of the queue
func (s *NodeQueue) Enqueue(t Vertex) {
	s.Lock.Lock()
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		s.Lock.Unlock()
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		if t.Distance < v.Distance {
			if k > 0 {
				s.Items = append(s.Items[:k+1], s.Items[k:]...)
				s.Items[k] = t
				insertFlag = true
			} else {
				s.Items = append([]Vertex{t}, s.Items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, t)
	}
	//s.items = append(s.items, t)
	s.Lock.Unlock()
}

// Dequeue removes an Node from the start of the queue
func (s *NodeQueue) Dequeue() *Vertex {
	s.Lock.Lock()
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	s.Lock.Unlock()
	return &item
}

// NewQ Creates New Queue
func (s *NodeQueue) NewQ() *NodeQueue {
	s.Lock.Lock()
	s.Items = []Vertex{}
	s.Lock.Unlock()
	return s
}

// IsEmpty returns true if the queue is empty
func (s *NodeQueue) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *NodeQueue) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items)
}

type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Edge
	Lock  sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.Lock.Lock()
	g.Nodes = append(g.Nodes, n)
	g.Lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node, weight int) {
	g.Lock.Lock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Edge)
	}
	ed1 := Edge{
		Node:   n2,
		Weight: weight,
	}

	ed2 := Edge{
		Node:   n1,
		Weight: weight,
	}
	g.Edges[*n1] = append(g.Edges[*n1], &ed1)
	g.Edges[*n2] = append(g.Edges[*n2], &ed2)
	g.Lock.Unlock()
}

func CreateGraph(data InputGraph) *ItemGraph {
	var g ItemGraph
	nodes := make(map[string]*Node)
	for _, v := range data.Graph {
		if _, found := nodes[v.Source]; !found {
			nA := Node{v.Source}
			nodes[v.Source] = &nA
			g.AddNode(&nA)
		}
		if _, found := nodes[v.Destination]; !found {
			nA := Node{v.Destination}
			nodes[v.Destination] = &nA
			g.AddNode(&nA)
		}
		g.AddEdge(nodes[v.Source], nodes[v.Destination], v.Weight)
	}
	return &g
}

func getShortestPath(startNode *Node, endNode *Node, g *ItemGraph) ([]string, int) {
	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)
	//pq := make(PriorityQueue, 1)
	//heap.Init(&pq)
	pq := NodeQueue{}
	//pq := q.NewQ()
	start := Vertex{
		Node:     startNode,
		Distance: 0,
	}
	for _, nval := range g.Nodes {
		dist[nval.Value] = math.MaxInt64
	}
	dist[startNode.Value] = start.Distance
	pq.Enqueue(start)
	//im := 0
	for !pq.IsEmpty() {
		v := pq.Dequeue()
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
					pq.Enqueue(store)
				}
				//visited[edge.Node.Value] = true
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

type Node struct {
	Value string
}

type Edge struct {
	Node   *Node
	Weight int
}

type Vertex struct {
	Node     *Node
	Distance int
}

// type ItemGraph struct {
// 	Nodes []*Node
// 	Edges map[Node][]*Edge
// 	Lock  sync.RWMutex
// }

type PriorityQueue []*Vertex

type NodeQueue struct {
	Items []Vertex
	Lock  sync.RWMutex
}

type InputGraph struct {
	Graph []InputData `json:"graph"`
	From  string      `json:"from"`
	To    string      `json:"to"`
}

type InputData struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
}
