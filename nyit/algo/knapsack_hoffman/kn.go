package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Fractional Knapsack
func fractionalKnapsack(values, weights []int, W int) float64 {
	type item struct {
		ratio  float64
		value  int
		weight int
	}

	n := len(weights)
	items := make([]item, n)
	for i := 0; i < n; i++ {
		items[i] = item{
			ratio:  float64(values[i]) / float64(weights[i]),
			value:  values[i],
			weight: weights[i],
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].ratio > items[j].ratio
	})

	totalVal := 0.0
	currentWeight := 0
	for _, it := range items {
		if currentWeight+it.weight <= W {
			totalVal += float64(it.value)
			currentWeight += it.weight
		} else {
			fraction := float64(W-currentWeight) / float64(it.weight)
			totalVal += fraction * float64(it.value)
			break
		}
	}
	return totalVal
}

// Huffman Coding
type Node struct {
	name      string
	frequency int
	left      *Node
	right     *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func hoffMan(frequencies map[string]int) *Node {
	pq := make(PriorityQueue, 0)
	for name, frequency := range frequencies {
		pq = append(pq, &Node{name: name, frequency: frequency})
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)
		merged := &Node{
			frequency: left.frequency + right.frequency,
			left:      left,
			right:     right,
		}
		heap.Push(&pq, merged)
	}

	return heap.Pop(&pq).(*Node)
}

func generateCodes(node *Node, code string, codes map[string]string) {
	if node == nil {
		return
	}
	if node.name != "" {
		codes[node.name] = code
		return
	}
	generateCodes(node.left, code+"0", codes)
	generateCodes(node.right, code+"1", codes)
}

func main() {
	values := []int{40, 50, 20}
	weights := []int{2, 5, 4}
	W := 6
	fmt.Printf("Fractional Knapsack Value: %.2f\n", fractionalKnapsack(values, weights, W))

	frequencies := map[string]int{"a": 5, "b": 9, "c": 12, "d": 13, "e": 16, "f": 45}
	root := hoffMan(frequencies)

	codes := make(map[string]string)
	generateCodes(root, "", codes)

	testStr := "abc"
	encoded := ""
	for _, char := range testStr {
		encoded += codes[string(char)]
	}

	fmt.Printf("Original string: %s\n", testStr)
	fmt.Printf("Huffman codes: %v\n", codes)
	fmt.Printf("Encoded string: %s\n", encoded)
}