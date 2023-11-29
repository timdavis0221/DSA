package stack

import "fmt"

// ImplementQueueUsingStacks implements a FIFO queue using two stacks at most.
func ImplementQueueUsingStacks() {
	queue := constructor()
	queue.push(1)
	queue.push(2)
	queue.peek()
	queue.pop()
	queue.empty()
}

type myQueue struct {
	stack []int
}

func constructor() *myQueue {
	fmt.Println("Initialized the instance of the queue")
	return &myQueue{
		stack: make([]int, 0),
	}
}

// push pushes element x to the back of the queue.
func (mq *myQueue) push(x int) {
	mq.stack = append(mq.stack, x)
}

// pop removes the element from the front of the queue and returns it.
func (mq *myQueue) pop() int {
	popValue := mq.stack[0]
	// Shifting all elements in the slice has a linear time complexity.
	// Because in the worst case, it involves copying each element to a new position in the slice.
	mq.stack = mq.stack[1:]
	return popValue
}

// peek returns the element at the front of the queue.
func (mq *myQueue) peek() int {
	return mq.stack[0]
}

// empty returns true if the queue is empty, false otherwise.
func (mq *myQueue) empty() bool {
	return len(mq.stack) == 0
}
