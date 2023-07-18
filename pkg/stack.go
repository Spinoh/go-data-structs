package day1

type Stack struct {
	Length int
	head   *StackNode
}

type StackNode struct {
	Value *int
	Prev  *StackNode
}

func (s *Stack) Push(val int) {
	node := &StackNode{
		Value: &val,
	}
	s.Length += 1

	if s.head == nil {
		s.head = node
		return
	}

	node.Prev = s.head
	s.head = node
}

func (s *Stack) Pop() *int {
	if s.Length == 0 {
		return nil
	}

	s.Length -= 1
	if s.Length == 0 {
		head := s.head
		s.head = nil

		return head.Value
	}

	head := s.head
	s.head = head.Prev

	return head.Value
}

func (s *Stack) Peek() *int {
	if s.head != nil {
		return s.head.Value
	}

	return nil
}
