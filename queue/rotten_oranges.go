package queue

import "fmt"

/*
Input:  arr[][C] = { {2, 1, 0, 2, 1},
                     {1, 0, 1, 2, 1},
                     {1, 0, 0, 2, 1}};
Output:
All oranges can become rotten in 2-time frames.
Explanation:
At 0th time frame:
 {2, 1, 0, 2, 1}
 {1, 0, 1, 2, 1}
 {1, 0, 0, 2, 1}

At 1st time frame:
 {2, 2, 0, 2, 2}
 {2, 0, 2, 2, 2}
 {1, 0, 0, 2, 2}

At 2nd time frame:
 {2, 2, 0, 2, 2}
 {2, 0, 2, 2, 2}
 {2, 0, 0, 2, 2}
*/

var row = []int{-1, 0, 0, 1}
var col = []int{0, -1, 1, 0}

type orangeLoc struct {
	x, y int
}

func findPassesToRotOranges(oranges [][]int) int {
	rotten := make([]*orangeLoc, 0)
	var passes int
	for i := 0; i < len(oranges); i++ {
		for j := 0; j < len(oranges[i]); j++ {
			if oranges[i][j] == 2 {
				rotten = append(rotten, &orangeLoc{
					x: i,
					y: j,
				})
			}
		}
	}
	for len(rotten) > 0 {
		currentQueue := rotten
		rotten = make([]*orangeLoc, 0)
		for _, orange := range currentQueue{
			for r:=0;r<len(row);r++{
				cr := orange.x+row[r]
				cy := orange.y + col[r]
				if isFreshAndValidPoint(cr,cy,oranges) {
					rotten = append(rotten, &orangeLoc{
						x: orange.x+row[r],
						y: orange.y + col[r],
					})
					oranges[cr][cy] = 2
				}
			}
		}
		passes++
		fmt.Println(oranges)
	}
	return passes
}

func isFreshAndValidPoint(x,y int, oranges [][]int) bool{
	if x<len(oranges) && x>=0 && y<len(oranges[x]) && y>=0 && oranges[x][y]==1{
		return true
	}
	return false
}

func PlayRotOranges() {
	oranges := [][]int{
		{2, 1, 0, 2, 1},
		{1, 0, 1, 2, 1},
		{1, 0, 0, 2, 1},
	}
	passes := findPassesToRotOranges(oranges)
	fmt.Println(passes)
}
