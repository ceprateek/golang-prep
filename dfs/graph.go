package dfs

import (
	"fmt"
	"math/rand"
)

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

func (g *Graph) PrintDFS(vertex int) {
	visited := make([]bool, g.vertices)
	g.traverseDFS(vertex, visited)
}

func (g *Graph) PrintCompleteDfs() {
	visited := make([]bool, g.vertices)
	for _, v := range rand.Perm(g.vertices) {
		g.traverseDFS(v, visited)
	}
}

func (g *Graph) traverseDFS(vertex int, visited []bool) {
	visited[vertex] = true
	fmt.Println(fmt.Sprintf("visiting: %d", vertex))
	for _, v := range g.adj[vertex] {
		if !visited[v] {
			g.traverseDFS(v, visited)
		}
	}
}

func Run() {
	g := Graph{}
	g.NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	g.PrintCompleteDfs()
}
