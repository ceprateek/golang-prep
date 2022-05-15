package algorithms

import "fmt"

type Tree struct {
	vertices int
	adjList  [][]int
}

func (t *Tree) InitTree(vertices int) {
	t.vertices = vertices
	t.adjList = make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		t.adjList[i] = make([]int, 0)
	}
}

func (t *Tree) AddEdge(source int, destination int) {
	t.adjList[source] = append(t.adjList[source], destination)
}

func (t *Tree) RunDfs(source int) {
	s := Stack{}
	visited := make([]bool, t.vertices)
	s.Push(source)
	visited[source] = true
	for !s.IsEmpty() {
		node := s.Pop()
		nodeVal := node.(int)
		fmt.Println(nodeVal)
		for _, neighborOfNode := range t.adjList[nodeVal] {
			if !visited[neighborOfNode] {
				s.Push(neighborOfNode)
			}
		}
	}
}

func (t *Tree) FindCommonParent(root int, nodes []int) int {
	q := make([]int, 1)
	parentList := make([]int, t.vertices)
	q[0] = root
	for len(q) > 0 {
		parent := q[0]
		for _, child := range t.adjList[parent] {
			parentList[child] = parent
			q = append(q, child)
		}
		q = q[1:]
	}
	cmp := make([]int, 0)
	for _, val := range nodes {
		if sliceContainsInt(cmp, parentList[val]){
			continue
		}
		cmp = append(cmp, parentList[val])
	}
	for len(cmp) > 1 {
		f := cmp[0]
		cmp = cmp[1:]
		parentOfF := parentList[f]
		if !sliceContainsInt(cmp, parentOfF){
			cmp = append(cmp, parentOfF)
		}
	}
	return cmp[0]
}

func sliceContainsInt(s []int, v int) bool{
	for _,val := range s{
		if val == v{
			return true
		}
	}
	return false
}

func PlayDfs() {
	t := &Tree{}
	t.InitTree(13)
	t.AddEdge(1, 2)
	t.AddEdge(1, 7)
	t.AddEdge(1, 8)
	t.AddEdge(2, 3)
	t.AddEdge(2, 6)
	t.AddEdge(3, 4)
	t.AddEdge(3, 5)
	t.AddEdge(8, 9)
	t.AddEdge(8, 12)
	t.AddEdge(9, 10)
	t.AddEdge(9, 11)
	t.RunDfs(1)
	cmp := t.FindCommonParent(1, []int{4,5,6,12})
	fmt.Printf("commond parent: %d\n", cmp)
}

type Stack []interface{}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(val interface{}) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}
