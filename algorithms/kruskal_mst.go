package algorithms

type DisjointSet struct {
	parent map[int]int
}

func (d *DisjointSet) InitDisjointSet(N int) {
	d.parent = make(map[int]int, N+1)
	for i := 0; i <= N; i++ {
		d.parent[i] = i
	}
}

func (d *DisjointSet) Find(i int) int {
	if d.parent[i] == i {
		return i
	} else {
		return d.Find(d.parent[i])
	}
}

//finding a cycle in the graph:
//for all edges check if there is an edge where both ends are already in disjoint set

//kruskal:
//sort all edges by weight
//if edge does not result in cycle add to result until
//len(result) = Vertices-1