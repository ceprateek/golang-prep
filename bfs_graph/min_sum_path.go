package bfs_graph

import (
	"fmt"
)

func PlayMinSumPath(){
	grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	minSum := minPathSum(grid)
	fmt.Println(minSum)
}

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i:=0;i<len(grid);i++{
		dp[i] = make([]int, len(grid[i]))
	}

	for i:=len(grid)-1;i>=0;i--{
		for j:=len(grid[i])-1;j>=0;j--{
			if i==len(grid)-1 && j==len(grid[i])-1{
				dp[i][j] = grid[i][j]
			}else if i==len(grid)-1 && j != len(grid[i])-1{
				dp[i][j] = grid[i][j] + dp[i][j+1]
			}else if j==len(grid[i])-1 && i!=len(grid)-1 {
				dp[i][j] = grid[i][j] + dp[i+1][j]
			}else{
				dp[i][j] = grid[i][j] + min(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]

}

func min(a, b int) int{
	if a<b{
		return a
	}
	return b
}
