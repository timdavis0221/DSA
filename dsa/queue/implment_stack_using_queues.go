package queue

import "fmt"

// ImplementStackUsingQueues implements a LIFO stack using two queues at most.
func ImplementStackUsingQueues() {
	stack := constructor()
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.top()
	stack.pop()
	stack.empty()
}

// myStack holds a slice as a queue to represent a stack.
type myStack struct {
	queue []int
}

// constructor initializes the instance of the stack.
func constructor() *myStack {
	fmt.Println("Create the instance of stack")
	myStack := &myStack{
		queue: make([]int, 0),
	}
	fmt.Printf("Initialized Stack: %v\n", myStack)
	return myStack
}

// push pushes element on the top of the stack.
func (ms *myStack) push(x int) {
	ms.queue = append(ms.queue, x)
	// Bring the last element to the front of queue to simulate the LIFO stack.
	for i := 0; i < len(ms.queue)-1; i++ {
		ms.queue = append(ms.queue, ms.queue[0])
		ms.queue = ms.queue[1:]
	}
	fmt.Printf("Stack.push(%d), stack: %v\n", x, ms.queue)
}

// pop removes the element on the top of the stack and returns it.
func (ms *myStack) pop() int {
	popValue := ms.queue[0]
	ms.queue = ms.queue[1:]
	fmt.Printf("Stack.pop(), stack: %v, pop value: %d\n", ms.queue, popValue)
	return popValue
}

// top returns the element on the top of the stack.
func (ms *myStack) top() int {
	fmt.Printf("Stack.top(), stack: %v, top value: %d\n", ms.queue, ms.queue[0])
	return ms.queue[0]
}

// empty returns true if the stack is empty, false otherwise.
func (ms *myStack) empty() bool {
	fmt.Printf("Stack.empty(): %t", len(ms.queue) == 0)
	return len(ms.queue) == 0
}
