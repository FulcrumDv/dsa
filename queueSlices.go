package main

import (
	"errors"
	"fmt"
)

// representation of a queue data struct
type Queue struct {
	Elements []int
}

// Enqueue adds an element of type int to the end of the queue
// q is the reciever variable
func (q *Queue) Enqueue(element int) {
	q.Elements = append(q.Elements, element)
}

// Dequeue will returns and removes the first element from the queue
func (q *Queue) Dequeue() (int, error) {
	// return first element
	// update the elements slice
	// validate that the queue is not empty
	if len(q.Elements) == 0 {
		return 0, errors.New("queue is empty")
	}

	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	return firstElement, nil

}

func main() {

	fmt.Println("Testing Queues")

	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue)

	queue.Dequeue()
	fmt.Println(queue)
}
