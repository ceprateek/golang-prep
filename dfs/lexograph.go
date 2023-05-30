package dfs

import (
	"fmt"
)

//https://leetcode.com/problems/lexicographically-smallest-equivalent-string/description/

func addNeighborToGraph(node byte, val byte, graph map[byte][]byte) {
	graph[node] = append(graph[node], val)
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	graph := make(map[byte][]byte)
	for i := 0; i < len(s1); i++ {
		addNeighborToGraph(s1[i], s2[i], graph)
		addNeighborToGraph(s2[i], s1[i], graph)
	}

	visited := make(map[byte]bool)
	minChar := make(map[byte]byte)

	var dfs func(node byte) byte
	dfs = func(node byte) byte {
		if _, ok := visited[node]; ok {
			return minChar[node]
		}
		visited[node] = true
		min := node
		for _, neighbor := range graph[node] {
			c := dfs(neighbor)
			if c < min {
				min = c
			}
		}
		minChar[node] = min
		return min
	}
	result := make([]byte, len(baseStr))
	for i := range baseStr {
		result[i] = dfs(baseStr[i])
	}
	return string(result)
}

func PlaySmallestEquivalentString() {
	fmt.Println(smallestEquivalentString("parker", "morris", "parser"))
}
