package bfs_graph

import "fmt"

type Graph struct {
	vertices int
	adj      [][]int
}

func (g *Graph) NewGraph(vertices int) {
	g.vertices = vertices
	g.adj = make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		g.adj[0] = make([]int, 0)
	}
}

func (g *Graph) AddEdge(v int, n int) {
	g.adj[v] = append(g.adj[v], n)
}

func (g *Graph) TraverseBFS(r int) {
	if r > g.vertices {
		return
	}
	visited := make([]bool, g.vertices)
	var queue []int
	queue = append(queue, r)
	visited[r] = true
	fmt.Println(r)
	for len(queue) > 0 {
		firstElement := queue[0]
		queue = queue[1:]
		for _, vertex := range g.adj[firstElement] {
			if !visited[vertex] {
				fmt.Println(vertex)
				visited[vertex] = true
				queue = append(queue, vertex)
			}
		}
	}
}

func Run() {
	g := &Graph{}
	g.NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	g.TraverseBFS(2)
}
