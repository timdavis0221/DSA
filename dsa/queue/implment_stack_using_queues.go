package queue

import "fmt"

// ImplementStackUsingQueues implements a LIFO stack using only two queues.
func ImplementStackUsingQueues() {
	stack := constructor()
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.top()
	stack.pop()
	stack.empty()
}

// myStack holds a slice as queue to represent a stack.
type myStack struct {
	intSlice []int
}

// constructor initializes the instance of the stack.
func constructor() *myStack {
	fmt.Println("Create the instance of stack")
	myStack := &myStack{
		intSlice: make([]int, 0),
	}
	fmt.Printf("Initialized Stack: %v\n", myStack)
	return myStack
}

// push pushes element on the top of the stack.
func (ms *myStack) push(x int) {
	ms.intSlice = append(ms.intSlice, x)
	// Bring the last element to the front of queue.
	for i := 0; i < len(ms.intSlice)-1; i++ {
		ms.intSlice = append(ms.intSlice, ms.intSlice[0])
		ms.intSlice = ms.intSlice[1:]
	}
	fmt.Printf("Stack.push(%d), stack: %v\n", x, ms.intSlice)
}

// pop removes the element on the top of the stack and returns it.
func (ms *myStack) pop() int {
	popValue := ms.intSlice[0]
	ms.intSlice = ms.intSlice[1:]
	fmt.Printf("Stack.pop(), stack: %v, pop value: %d\n", ms.intSlice, popValue)
	return popValue
}

// top returns the element on the top of the stack.
func (ms *myStack) top() int {
	fmt.Printf("Stack.top(), stack: %v, top value: %d\n", ms.intSlice, ms.intSlice[0])
	return ms.intSlice[0]
}

// empty returns true if the stack is empty, false otherwise.
func (ms *myStack) empty() bool {
	fmt.Printf("Stack.empty(): %t", len(ms.intSlice) == 0)
	return len(ms.intSlice) == 0
}
