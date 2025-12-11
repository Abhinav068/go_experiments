package main

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	first := q.items[0]
	q.items = q.items[1:] // pop from front
	return first, true
}

func (q *Queue[T]) Front() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	return q.items[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// q.items = q.items[1:] is O(1) but keeps the underlying array in memory and can cause memory leaks for very large queues.
// https://chatgpt.com/share/693a474c-5484-8012-baee-329c250dcc13