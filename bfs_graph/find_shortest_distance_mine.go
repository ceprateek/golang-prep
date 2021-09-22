package bfs_graph

import "fmt"

// 0 x 0 0 0
// 0 0 M 0 0
// 0 0 0 0 0
// 0 x 0 0 0

// 0 x 1 0 0
// 0 1 M 1 0
// 0 2 1 2 0
// 0 x 1 0 0

type maze struct {
	maze [][]uint8
}

func (m *maze) NewMaze() {
	m.maze = [][]uint8{
		{'O', 'M', 'O', 'O', 'X'},
		{'O', 'X', 'X', 'O', 'M'},
		{'O', 'O', 'O', 'O', 'O'},
		{'O', 'X', 'X', 'X', 'O'},
		{'O', 'O', 'M', 'O', 'O'},
		{'O', 'X', 'X', 'M', 'O'},
	}
}

func (m *maze) GetShortestDistance() (result [][]int) {
	result = make([][]int, len(m.maze))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(m.maze[i]))
	}

	for i := 0; i < len(m.maze); i++ {
		for j := 0; j < len(m.maze[i]); j++ {
			if m.maze[i][j] == 'M' {
				result[i][j] = 0
			}else {
				result[i][j] = -1
			}
		}
	}

	var queue []*point
	for i := 0; i < len(m.maze); i++ {
		for j := 0; j < len(m.maze[i]); j++ {
			if m.maze[i][j] == 'M' {
				queue = append(queue, &point{x: i, y: j})
			}
		}
	}
	var temp []*point
	for len(queue) > 0 {
		temp = queue[:]
		queue = make([]*point, 0)
		for _, p := range temp {
			var pointsToCheck []*point
			pointsToCheck = append(pointsToCheck, &point{x: p.x - 1, y: p.y}, &point{x: p.x + 1, y: p.y},
				&point{x: p.x, y: p.y - 1}, &point{x: p.x, y: p.y + 1})
			for _, pointToCheck := range pointsToCheck {
				if m.isValidPoint(pointToCheck) && m.maze[pointToCheck.x][pointToCheck.y] != 'X' &&
					m.maze[pointToCheck.x][pointToCheck.y] != 'M' {
					current := result[pointToCheck.x][pointToCheck.y]
					newval := result[p.x][p.y] + 1
					if current == -1 {
						result[pointToCheck.x][pointToCheck.y] = newval
						queue = append(queue, pointToCheck)
					} else if (newval < current) {
						result[pointToCheck.x][pointToCheck.y] = newval
						queue = append(queue, pointToCheck)
					} else {
						//already contains the shorter path
					}
				}
			}
		}
	}

	return result
}

func (a *maze) isValidPoint(p *point) bool {
	if p.y >= 0 && p.y < len(a.maze[0]) && p.x >= 0 && p.x < len(a.maze) {
		return true
	}
	return false
}

type point struct {
	x, y int
}

func RunMaze() {
	m := &maze{}
	m.NewMaze()
	fmt.Println(m.GetShortestDistance())
}
