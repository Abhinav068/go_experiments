package optimized

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	n := len(s.items)
	if n == 0 {
		return zero, false
	}
	v := s.items[n-1]

	// Prevent memory retention
	s.items[n-1] = zero

	s.items = s.items[:n-1]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int { return len(s.items) }

func (s *Stack[T]) IsEmpty() bool { return len(s.items) == 0 }

func (s *Stack[T]) Clear() {
	var zero T
	for i := range s.items {
		s.items[i] = zero
	}
	s.items = s.items[:0]
}

// https://chatgpt.com/share/693a474c-5484-8012-baee-329c250dcc13