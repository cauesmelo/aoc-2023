package stack

type node[T any] struct {
	value T
	prev  *node[T]
}

type Stack[T any] struct {
	top *node[T]
	len int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(c T) {
	n := &node[T]{
		value: c,
		prev:  s.top,
	}

	s.top = n
	s.len++
}

func (s *Stack[T]) Pop() T {
	top := s.top

	s.len--
	s.top = top.prev

	return top.value
}

func (s *Stack[T]) Len() int {
	return s.len
}
