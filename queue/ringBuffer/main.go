package main

import "fmt"

// RingQueue is a circular/ring buffer queue with automatic growth.
// Amortized O(1) Enqueue, O(1) Dequeue.
type RingQueue[T any] struct {
	data     []T
	head     int // index of first element (if size>0)
	tail     int // index for next write (i.e., (head+size) % cap)
	size     int
	capacity int
}

// NewRingQueue creates a queue with given initial capacity (min 1).
func NewRingQueue[T any](initialCap int) *RingQueue[T] {
	if initialCap <= 0 {
		initialCap = 8
	}
	return &RingQueue[T]{
		data:     make([]T, initialCap),
		capacity: initialCap,
	}
}

// Len returns number of items in queue.
func (q *RingQueue[T]) Len() int { return q.size }

// Cap returns current capacity of the queue.
func (q *RingQueue[T]) Cap() int { return q.capacity }

// IsEmpty returns true when queue has no items.
func (q *RingQueue[T]) IsEmpty() bool { return q.size == 0 }

// Enqueue adds v to the queue. Grows if full.
func (q *RingQueue[T]) Enqueue(v T) {
	if q.size == q.capacity {
		q.grow()
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	q.size++
}

// Dequeue removes and returns the front element. Second return is false when empty.
func (q *RingQueue[T]) Dequeue() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	val := q.data[q.head]
	// Clear slot to avoid memory retention for pointer types
	q.data[q.head] = zero
	q.head = (q.head + 1) % q.capacity
	q.size--
	return val, true
}

// Peek returns the front element without removing. Second return is false when empty.
func (q *RingQueue[T]) Peek() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	return q.data[q.head], true
}

// Clear empties the queue and zeroes stored elements.
func (q *RingQueue[T]) Clear() {
	if q.size == 0 {
		return
	}
	// zero out all used slots
	for i := 0; i < q.size; i++ {
		idx := (q.head + i) % q.capacity
		var zero T
		q.data[idx] = zero
	}
	q.head, q.tail, q.size = 0, 0, 0
}

// grow doubles the capacity and reorders elements starting at index 0.
func (q *RingQueue[T]) grow() {
	newCap := q.capacity * 2
	if newCap == 0 {
		newCap = 8
	}
	newData := make([]T, newCap)

	// copy in-order elements
	if q.size > 0 {
		if q.head < q.tail {
			copy(newData, q.data[q.head:q.tail])
		} else {
			// wrapped
			n := copy(newData, q.data[q.head:q.capacity])
			copy(newData[n:], q.data[0:q.tail])
		}
	}
	q.data = newData
	q.head = 0
	q.tail = q.size
	q.capacity = newCap
}

func main() {
	q := NewRingQueue[int](5)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if v, ok := q.Dequeue(); ok {
		fmt.Println(v) // 10
	}

	q.Enqueue(40)
	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		fmt.Println(v) // 20 30 40
	}
}
