package bfs_graph

import (
	"fmt"
	"math"
	"sort"
)

func PlayDijkstra() {
	var edges []*weighedEdge
	edges = append(edges, &weighedEdge{0, 1, 10})
	edges = append(edges, &weighedEdge{0, 4, 3})
	edges = append(edges, &weighedEdge{1, 2, 2})
	edges = append(edges, &weighedEdge{1, 4, 4})
	edges = append(edges, &weighedEdge{2, 3, 9})
	edges = append(edges, &weighedEdge{3, 2, 7})
	edges = append(edges, &weighedEdge{4, 1, 1})
	edges = append(edges, &weighedEdge{4, 2, 8})
	edges = append(edges, &weighedEdge{4, 3, 2})
	N := 5
	graph := weighedGraph{}
	graph.newWeighedGraph(edges, N)
}

type weighedEdge struct {
	source, destination, weight int
}

type weighedGraph struct {
	adjList [][]*weighedEdge
}

type Node struct {
	vertex int
	weight int
}

func (w *weighedGraph) newWeighedGraph(edges []*weighedEdge, numberOfEdges int) {
	w.adjList = make([][]*weighedEdge, numberOfEdges)
	for edge := 0; edge < numberOfEdges; edge++ {
		w.adjList[edge] = make([]*weighedEdge, 0)
	}
	for _, edge := range edges {
		w.adjList[edge.source] = append(w.adjList[edge.source], edge)
	}
}

func (w *weighedGraph) findShortestPath2Vertices(source int, destination int, N int) {
	minHeap := &PriorityQueue{}
	minHeap.Push(&Node{
		vertex: source,
		weight: 0,
	})
	distance := make([]int, N)
	done := make([]bool, N)
	previous := make([]int, N)
	for i := 0; i < N; i++ {
		distance[i] = math.MaxInt32
	}
	distance[source] = 0
	previous[source] = -1
	for minHeap.Len() > 0 && !done[destination] {
		node := minHeap.Poll()
		ver := node.vertex
		for _, neighbor := range w.adjList[ver] {
			currentNeighborDestination := neighbor.destination
			weightNeighborDestination := neighbor.weight
			if !done[currentNeighborDestination] && distance[currentNeighborDestination] > distance[ver]+weightNeighborDestination {
				distance[currentNeighborDestination] = distance[ver] + weightNeighborDestination
				previous[currentNeighborDestination] = ver
				minHeap.Push(&Node{
					vertex: currentNeighborDestination,
					weight: distance[currentNeighborDestination],
				})
			}
		}
		done[ver] = true
	}
	printPath(previous, source, destination)
}

func printPath(previous []int, source int, destination int) {
	if destination > 0 {
		printPath(previous, source, previous[destination])
		fmt.Println(destination)
	} else {
		fmt.Println(source)
	}
}

func (w *weighedGraph) findShortestPath(source int, N int) {
	minHeap := &PriorityQueue{}
	minHeap.Push(&Node{
		vertex: source,
		weight: 0,
	})
	distance := make([]int, N)
	done := make([]bool, N)
	previousEdge := make([]int, N)
	for i := 0; i < N; i++ {
		distance[i] = math.MaxInt32
		done[i] = false
	}
	distance[source] = 0
	done[source] = true
	previousEdge[source] = -1

	for minHeap.Len() > 0 {
		node := minHeap.Poll()
		ver := node.vertex
		fmt.Println(fmt.Sprintf("processing: %d", ver))
		for _, neighbor := range w.adjList[ver] {
			toProcess := neighbor.destination
			weight := neighbor.weight
			if !done[toProcess] && (distance[ver]+weight < distance[toProcess]) {
				distance[toProcess] = distance[ver] + weight
				previousEdge[toProcess] = ver
				minHeap.Push(&Node{
					vertex: toProcess,
					weight: distance[toProcess],
				})
			}
		}
		done[ver] = true
	}
	for i := 0; i < N; i++ {
		if i != source {
			sroute := make([]int, 0)
			getRoute(&sroute, previousEdge, i)
			fmt.Println(fmt.Sprintf("(%d -> %d) min: %d route: %v", source, i, distance[i], sroute))
		}
	}
}

func getRoute(route *[]int, prev []int, edge int) {
	if edge >= 0 {
		getRoute(route, prev, prev[edge])
		*route = append(*route, edge)
	}
}

type PriorityQueue []*Node

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p *PriorityQueue) Push(node *Node) {
	*p = append(*p, node)
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].weight < (*p)[j].weight
	})
}

func (p *PriorityQueue) Push2(node *Node) {
	*p = append(*p, node)
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].weight > (*p)[j].weight
	})
}

func (p *PriorityQueue) Poll() *Node {
	i := (*p)[0]
	*p = (*p)[1:]
	return i
}
