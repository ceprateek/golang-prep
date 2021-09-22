//Find minimum passes required to convert all negative values in a matrix
//https://www.techiedelight.com/find-minimum-passes-required-convert-negative-values-matrix/

package bfs_graph

import "fmt"

type Matrix struct {
	Edges [][]int
}

func (m *Matrix) NewMatrix(row int, col int) {
	m.Edges = make([][]int, row)
	for i := 0; i < row; i++ {
		m.Edges[i] = make([]int, col)
	}
}

func (m *Matrix) InsertVal(row, col, val int) {
	m.Edges[row][col] = val
}

func RunMatrix() {
	m := &Matrix{}
	m.NewMatrix(4, 5)
	m.InsertVal(0, 0, -1)
	m.InsertVal(0, 1, -9)
	m.InsertVal(0, 3, -1)
	m.InsertVal(1, 0, -8)
	m.InsertVal(1, 1, -3)
	m.InsertVal(1, 2, -2)
	m.InsertVal(1, 3, 9)
	m.InsertVal(1, 4, -7)
	m.InsertVal(2, 0, 2)
	m.InsertVal(2, 3, -6)
	m.InsertVal(3, 1, -7)
	m.InsertVal(3, 2, -3)
	m.InsertVal(3, 3, 5)
	m.InsertVal(3, 4, -4)
	fmt.Println(m.Edges)
	fmt.Println(m.convertNegatives())
}

type position struct {
	x, y int
}

func (m *Matrix) convertNegatives() int {
	var allPositives []*position
	for i:=0; i< len(m.Edges); i++ {
		for j:=0; j<len(m.Edges[0]); j++ {
			if m.Edges[i][j] > 0 {
				p := position{
					x: i,
					y: j,
				}
				allPositives = append(allPositives, &p)
			}
		}
	}
	var pass int
	for len(allPositives) > 0 {
		currentQueue := allPositives[:]
		allPositives = make([]*position, 0)
		for len(currentQueue) > 0 {
			firstPoint := currentQueue[0]
			currentQueue = currentQueue[1:]
			pointsToCheck := make([]position, 0)
			pointsToCheck = append(pointsToCheck, position{x: firstPoint.x - 1, y: firstPoint.y})
			pointsToCheck = append(pointsToCheck, position{x: firstPoint.x + 1, y: firstPoint.y})
			pointsToCheck = append(pointsToCheck, position{x: firstPoint.x, y: firstPoint.y - 1})
			pointsToCheck = append(pointsToCheck, position{x: firstPoint.x, y: firstPoint.y + 1})
			for _, point := range pointsToCheck {
				if m.isValidPoint(&point) && m.Edges[point.x][point.y] < 0 {
					m.Edges[point.x][point.y] = -m.Edges[point.x][point.y]
					allPositives = append(allPositives, &point)
					fmt.Println(m.Edges[point.x][point.y])
				}
			}
		}
		pass++
		fmt.Println(m.Edges)
	}
	return pass
}

func (a *Matrix) isValidPoint(point *position) bool {
	if point.y >= 0 && point.y < len(a.Edges[0]) && point.x >= 0 && point.x < len(a.Edges) {
		return true
	}
	return false
}
