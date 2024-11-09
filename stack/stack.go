package stack

type StackItem struct {
	item interface{}
	next *StackItem
}

// Stack is a base structure for LIFO
type Stack struct {
	sp    *StackItem
	depth uint64
}

type Stack1 struct {
	sp    *StackItem
	depth uint32
}

// Initialzes new Stack
func New() *Stack {
	var stack *Stack = new(Stack)

	stack.depth = 0
	return stack
}
func New1() *Stack1 {
	stack := new(Stack1)
	stack.depth = 0
	return stack
}

// Pushes a given item into Stack
func (stack *Stack) Push(item interface{}) {
	stack.sp = &StackItem{item: item, next: stack.sp}
	stack.depth++
}

func (s1 *Stack1) Push(item interface{}) {
	s1.sp = &StackItem{item, s1.sp}
	s1.depth++
}

// Deletes top of a stack and return it
func (stack *Stack) Pop() interface{} {
	if stack.depth > 0 {
		item := stack.sp.item
		stack.sp = stack.sp.next
		stack.depth--
		return item
	}

	return nil
}

func (s1 *Stack1) Pop() interface{} {
	if s1.depth == 0 {
		return nil
	}
	item := s1.sp.item
	s1.sp = s1.sp.next
	s1.depth--
	return item
}

// Peek returns top of a stack without deletion
func (stack *Stack) Peek() interface{} {
	if stack.depth > 0 {
		return stack.sp.item
	}

	return nil
}
func (s1 *Stack1) Peek() interface{} {
	if s1.depth == 0 {
		return nil
	}
	return s1.sp.item
}
