package stacks

type Stack []interface{}

func (s *Stack) Pop() interface{} {
	if len(*s) == 0{
		return nil
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *Stack) Clear() {
	for !s.IsEmpty(){
		s.Pop()
	}
}

func (s *Stack) Push(val interface{}) {
	*s = append(*s, val)
}

func (s *Stack) Peek() interface{} {
	if len(*s) == 0{
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0{
		return true
	}
	return false
}