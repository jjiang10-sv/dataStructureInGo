package graph

//: If there is a path from u to v, then v appears after u in the ordering
type Graph1 struct {
	edges    map[int][]int
	vertices []int
}

func (g *Graph1) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
}

func (g *Graph1) topologicalSortUtil(v int, visited map[int]bool, stack *[]int) {
	visited[v] = true

	for _, u := range g.edges[v] {
		if !visited[u] {
			g.topologicalSortUtil(u, visited, stack)
		}
	}

	*stack = append([]int{v}, *stack...)
}

func (g *Graph1) topologicalSort() []int {
	stack := []int{}
	visited := make(map[int]bool)

	for _, v := range g.vertices {
		if !visited[v] {
			g.topologicalSortUtil(v, visited, &stack)
		}
	}

	return stack
}
