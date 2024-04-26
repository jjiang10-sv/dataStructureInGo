package prac

type stackItem struct {
	item interface{}
	next *stackItem
}

type stack struct {
	depth uint
	top *stackItem
}

func newStack() *stack {
	return new(stack)
}

func (s *stack) push(item interface{}) {
	data := &stackItem{item: item, next: s.top}
	s.top = data
	s.depth++
}
func (s *stack) pop() interface{} {
	if s.depth == 0 {
		return nil
	}
	data := s.top
	s.top = data.next
	s.depth--
	return data.item

}
func(s *stack) peek() interface{} {
	return s.top
}