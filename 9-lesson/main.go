// Queue data structure

package main

import "fmt"

func main() {
	testQueue()
}

func testQueue() {
	queue := NewQueue()
	queue.Enqueue(100)
	fmt.Println("queue: ", queue)
	queue.Enqueue(25).Enqueue(77)
	fmt.Println("queue: ", queue)
	fmt.Println("is queue empty? ", queue.IsEmpty())

	var result, _ = queue.Peek()
	fmt.Println("the next item in the queue: ", result)

	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println("Dequeue one item...")
	result, _ = queue.Dequeue()
	fmt.Println(result)
	fmt.Println(queue.IsEmpty())
	_, err := queue.Peek()
	fmt.Println(err)
}

// declare the Queue type
type Queue struct {
	data []int
}

// NewQueue : returns a new Queue object
func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// IsEmpty : checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek : Returns the next item in the queue
func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

// Enqueue : Adds a new element to the queue and returns a pointer to the queue
func (q *Queue) Enqueue(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

// Dequeue : deletes the next element from the queue and returns its value
func (q *Queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	element := q.data[0]
	q.data = q.data[1:] // returns all elements starting with the first element
	return element, nil
}
