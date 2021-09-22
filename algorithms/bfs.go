package algorithms

import "fmt"

type Graph struct {
	vertices int
	adjMap [][]int
}

func (g *Graph) AddEdge(vertice int, neighbor int)  {
	g.adjMap[vertice] = append(g.adjMap[vertice], neighbor)
}

func (g *Graph) NewGraph(vertices int) {
	g.vertices = vertices
	g.adjMap = make([][]int, vertices)
	for i:=0;i<vertices;i++{
		g.adjMap[i] = make([]int, 0)
	}
}

func (g *Graph) TraverseBFS(source int) {
	visited := make([]bool, g.vertices)
	var queue []int
	queue = append(queue, source)
	visited[source] = true
	for len(queue) > 0 {
		firstElementFromQueue := queue[0]
		queue = queue[1:]
		allNeighbors := g.adjMap[firstElementFromQueue]
		for _, neighbor := range allNeighbors {
			if !visited[neighbor]{
				fmt.Println(neighbor)
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func PlayBfs() {
	g := &Graph{}
	g.NewGraph(6)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.TraverseBFS(2)
}
