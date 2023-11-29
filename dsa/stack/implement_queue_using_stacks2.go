package stack

import (
	"fmt"
)

// ImplementQueueUsingStacks2 implements a FIFO queue using two stacks.
func ImplementQueueUsingStacks2() {
	queue := new()
	queue.push(1)
	queue.push(2)
	queue.peek()
	queue.pop()
	queue.empty()
}

type queue struct {
	enqueueStack, dequeueStack []int
}

type emptyError struct {
	msg string
}

func (e emptyError) Error() string {
	return fmt.Sprintf(e.msg)
}

func new() *queue {
	fmt.Println("Initialized the instance of the queue")
	return &queue{
		enqueueStack: make([]int, 0),
		dequeueStack: make([]int, 0),
	}
}

// push pushes element x to the back of the queue.
func (mq *queue) push(x int) {
	mq.enqueueStack = append(mq.enqueueStack, x)
}

// transferToDequeueStack performs transfer of elements from LIFO enqueueStack to LIFO dequeueStack
// for amortizing the execution time of pop() and peek() in O(1).
// O(1) if dequeueStack is not empty, O(n) otherwise.
func (mq *queue) transferToDequeueStack() error {
	if len(mq.dequeueStack) == 0 {
		for len(mq.enqueueStack) > 0 {
			element := mq.enqueueStack[len(mq.enqueueStack)-1]
			mq.dequeueStack = append(mq.dequeueStack, element)
			mq.enqueueStack = mq.enqueueStack[:len(mq.enqueueStack)-1]
		}
	}
	if len(mq.dequeueStack) == 0 {
		return emptyError{
			msg: "dequeueStack is empty",
		}
	}
	return nil
}

// pop removes the element from the front of the queue and returns it.
func (mq *queue) pop() int {
	if err := mq.transferToDequeueStack(); err != nil {
		return -1
	}
	popValue := mq.dequeueStack[len(mq.dequeueStack)-1]
	mq.dequeueStack = mq.dequeueStack[:len(mq.dequeueStack)-1]
	return popValue
}

// peek returns the element at the front of the queue.
func (mq *queue) peek() int {
	if err := mq.transferToDequeueStack(); err != nil {
		return -1
	}
	return mq.dequeueStack[len(mq.dequeueStack)-1]
}

// empty returns true if the queue is empty, false otherwise.
func (mq *queue) empty() bool {
	return (len(mq.enqueueStack) == 0 && len(mq.dequeueStack) == 0)
}
