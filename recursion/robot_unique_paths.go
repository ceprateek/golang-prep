package recursion

import "fmt"

func uniquePaths(m int, n int) int {
	routes := 0
	var paths []string
	var s string
	uniquePathHelper(0, 0, m, n, &routes, &paths, &s)
	for _, path := range paths{
		fmt.Println(path)
	}

	return routes
}

func uniquePathHelper(y, x, m, n int, r *int, paths *[]string, pathsofar *string) {
	if y == m-1 && x == n-1 {
		*r++
		temp := *pathsofar
		*paths = append(*paths, temp)
		return
	}
	if y >= m || x >= n {
		return
	}
	*pathsofar = *pathsofar + "down"
	uniquePathHelper(y+1, x, m, n, r,paths,pathsofar)
	*pathsofar = *pathsofar + "right"
	uniquePathHelper(y, x+1, m, n, r,paths,pathsofar)
}

func PlayRobotPaths() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
	//fmt.Println(findFibnocci(10))
}

func findFibnocci(n int) int{
	return fibHelper(n)
}

func fibHelper(n int) int{
	if n <= 1{
		return n
	}else {
		return fibHelper(n-1)+fibHelper(n-2)
	}
}