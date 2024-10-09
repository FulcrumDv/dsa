package main

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

// will add element to front of the queue
func (q *Queue) Enqueue(element int) {
	newNode := &Node{value: element}
	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

// dequeue will return and move the first element in queue
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	element := q.front.value
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return element, true
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) getString() {
	current := q.front
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}

}

func main() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue getString()

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	fmt.Println(queue.IsEmpty())

}
