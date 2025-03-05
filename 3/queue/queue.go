package main

import "fmt"

type QueueNode[T any] struct {
	data T
	next *QueueNode[T]
}

type Queue[T any] struct {
	last  *QueueNode[T]
	first *QueueNode[T]
}

func (q *Queue[T]) add(data T) {
	newNode := &QueueNode[T]{data: data}

	if q.last != nil {
		q.last.next = newNode
	} else {
		q.first = newNode // First element case
	}

	q.last = newNode
}

func (q *Queue[T]) remove() {
	if q.first == nil {
		return // Queue is empty
	}
	q.first = q.first.next
	if q.first == nil {
		q.last = nil // Queue is now empty
	}

}
func (q *Queue[T]) peek() (T, bool) {
	if q.first == nil {
		var zero T
		return zero, false
	}
	return q.first.data, true
}

func (q *Queue[T]) isEmpty() bool {
	if q.first == nil {
		return true
	}
	return false

}

func main() {
	queue := Queue[int]{}
	queue.add(1)
	queue.add(5)
	queue.add(3)

	peeked, exists := queue.peek()
	if exists {
		fmt.Println("Front element:", peeked)
	} else {
		fmt.Println("Queue is empty")
	}

	queue.remove()

	peeked, exists = queue.peek()
	if exists {
		fmt.Println("Front element after remove:", peeked)
	} else {
		fmt.Println("Queue is empty")
	}

}
