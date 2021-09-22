package bfs_graph

import "fmt"

var row = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var col = []int{-1, 0, 1, -1, 1, -1, 0, 1}

type connectedIslandMatrix struct {
	matrix  [][]int
	visited [][]bool
}

func (m *connectedIslandMatrix) newMatrix() {
	m.matrix = [][]int{
		{1, 0, 1, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 1, 0, 1, 0, 1, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 1, 0, 0, 0},
		{1, 0, 0, 1, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 1, 1, 1},
		{0, 1, 0, 1, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 1, 1, 0},
		{1, 0, 1, 0, 1, 0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 1, 1, 1},
	}
	m.visited = make([][]bool, len(m.matrix))
	for i := 0; i < len(m.matrix); i++ {
		m.visited[i] = make([]bool, len(m.matrix[i]))
	}
}

func (m *connectedIslandMatrix) findIslands(p *point) {
	queue := make([]*point, 0)
	queue = append(queue, p)
	m.visited[p.x][p.y] = true
	for len(queue) > 0 {
		//get first element from queue and remove the same from queue
		firstElement := queue[0]
		queue = queue[1:]
		for r := 0; r < len(row); r++ {
			neighbor := &point{
				x: firstElement.x + row[r],
				y: firstElement.y + col[r],
			}
			if m.isPointSafe(neighbor) {
				m.visited[neighbor.x][neighbor.y] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func (m *connectedIslandMatrix) isPointSafe(p *point) bool {
	//checks if the point is valid index on the matrix
	//also makes sure the point has not been processed
	if p.x >= 0 && p.x < len(m.matrix) && p.y >= 0 &&
		p.y < len(m.matrix[p.x]) && !m.visited[p.x][p.y] && m.matrix[p.x][p.y] == 1 {
		return true
	}
	return false
}

func PlayConnectedIslands() {
	m := &connectedIslandMatrix{}
	m.newMatrix()
	var islands int
	for r:=0;r< len(m.matrix);r++{
		for c:=0; c< len(m.matrix[r]); c++{
			if m.matrix[r][c] == 1 && !m.visited[r][c]{
				p := &point{x: r, y: c,}
				m.findIslands(p)
				islands++
			}
		}
	}
	fmt.Println(islands)
}
