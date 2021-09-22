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
	for !s.IsEmpty(){
		node := s.Pop()
		nodeVal := node.(int)
		fmt.Println(nodeVal)
		for _, neighborOfNode := range t.adjList[nodeVal]{
			if !visited[neighborOfNode] {
				s.Push(neighborOfNode)
			}
		}
	}
}

func PlayDfs(){
	t := &Tree{}
	t.InitTree(13)
	t.AddEdge(1,2)
	t.AddEdge(1,7)
	t.AddEdge(1,8)
	t.AddEdge(2,3)
	t.AddEdge(2,6)
	t.AddEdge(3,4)
	t.AddEdge(3,5)
	t.AddEdge(8,9)
	t.AddEdge(8,12)
	t.AddEdge(9,10)
	t.AddEdge(9,11)
	t.RunDfs(1)
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
func (s *Stack) Pop() (interface{}) {
	if s.IsEmpty() {
		return nil
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element
	}
}
