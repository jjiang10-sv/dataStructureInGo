package graph

import (
	"fmt"
	"testing"
)

func compare(x interface{}, y interface{}) bool {
	return x.(int) < y.(int)
}

func Test_topsort1(t *testing.T) {
	g := &Graph1{
		edges:    make(map[int][]int),
		vertices: []int{5, 4, 2, 3, 1, 0},
	}

	g.addEdge(5, 2)
	g.addEdge(5, 0)
	g.addEdge(4, 0)
	g.addEdge(4, 1)
	g.addEdge(2, 3)
	g.addEdge(3, 1)

	result := g.topologicalSort()

	for _, v := range result {
		fmt.Println(v)
	}
}

func Test_shortestPath(t *testing.T) {
	inputGraph := InputGraph{
		Graph: []InputData{
			{
				Source:      "A",
				Destination: "B",
				Weight:      1,
			},
			{
				Source:      "A",
				Destination: "C",
				Weight:      2,
			},
			{
				Source:      "B",
				Destination: "D",
				Weight:      3,
			},
			{
				Source:      "B",
				Destination: "E",
				Weight:      5,
			},
			{
				Source:      "C",
				Destination: "D",
				Weight:      4,
			},
			{
				Source:      "D",
				Destination: "E",
				Weight:      2,
			},
			{
				Source:      "E",
				Destination: "F",
				Weight:      1,
			},
			{
				Source:      "C",
				Destination: "F",
				Weight:      3,
			},
		},
		From: "A",
		To:   "E",
	}

	itemGraph := CreateGraph(inputGraph)

	//getShortestPath(&Node{"A"}, &Node{"E"}, itemGraph)
	getShortestPath1(&Node{"A"}, &Node{"E"}, itemGraph)

}

func TestHouseRobDp(t *testing.T) {
	input := []uint{1, 2, 3, 1, 1, 2, 3, 1}
	res := robHouseDp(input)
	println(res)

}

func TestHoffmanCodes(t *testing.T) {
	vals := []string{"j", "o", "h", "d", "n"}
	frequencies := []uint{5, 2, 8, 2, 7}
	res := map[string]string{}
	hoffman_code_tree(vals, frequencies).getValCodefromHoffmanTree(res, "")
	for item, code := range res {
		println(item, code)
	}
	coded := NewHoffmanCodes(res).code("john")
	println(coded)

}

func TestAllPairShortestPath(t *testing.T) {
	allPairShortestPath()

}
