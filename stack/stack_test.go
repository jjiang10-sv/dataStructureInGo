package stack

import "testing"

func TestStack(t *testing.T) {
	var stack *Stack = New()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for i := 5; i > 0; i-- {
		item := stack.Pop()

		if item != i {
			t.Error("TestStack failed...", i)
		}
	}
}
func TestStack1(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	s1 := New1()
	for _, item := range arr1 {
		s1.Push(item)
	}
	for i := 5; i > 0; i-- {
		item := s1.Pop()
		if item != i {
			t.Error("test stack failed", i)
		}
	}
}
