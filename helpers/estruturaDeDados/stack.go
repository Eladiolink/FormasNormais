package estruturadedados

type Node struct {
	Data interface{}
	Next *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(data interface{}) {
	newNode := &Node{Data: data, Next: s.Top}
	s.Top = newNode
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.Top == nil {
		return nil, false
	}
	data := s.Top.Data
	s.Top = s.Top.Next
	return data, true
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}
