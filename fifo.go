package queue

// FIFO is a naive implementation of a Generic FIFO queue
type FIFO[T any] struct {
	data []T
}

// Next gets the next element in the queue and removes it from the Queue
func (q *FIFO[T]) Next() T {
	n := q.data[0]
	q.data = append(make([]T, 0, len(q.data[1:])), q.data[1:]...)
	return n
}

// More returns if Queue has more data to give
func (q *FIFO[T]) More() bool {
	return len(q.data) != 0
}

// Add an element to the Queue
func (q *FIFO[T]) Add(t T) {
	q.data = append(q.data, t)
}

// Size of the Queue
func (q *FIFO[T]) Size() int {
	return len(q.data)
}

// All elements in the queue as a copy of the current queue
func (q *FIFO[T]) All() []T {
	return append(make([]T, 0, len(q.data)), q.data...)
}

// Clear Queue back to initial state
func (q *FIFO[T]) Clear() {
	q.data = *new([]T)
}
