package queue

// Queue is a naive implementation of a FIFO genric queue
type Queue[T any] struct {
	data []T
}

// Next gets the next element in the queue and removes it from the Queue
func (q *Queue[T]) Next() T {
	n := q.data[0]
	var newData []T
	newData = append(newData, q.data[1:]...)
	q.data = newData
	return n
}

// More returns if Queue has more data to give
func (q *Queue[T]) More() bool {
	return len(q.data) != 0
}

// Add an element to the Queue
func (q *Queue[T]) Add(t T) {
	q.data = append(q.data, t)
}

// Size of the Queue
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// All elements in the queue as a copy of the current queue
func (q *Queue[T]) All() []T {
	var all []T
	all = append(all, q.data...)
	return all
}

// Clear Queue back to initial state
func (q *Queue[T]) Clear() {
	q.data = *new([]T)
}
